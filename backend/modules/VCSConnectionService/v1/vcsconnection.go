package VCSConnectionService

import (
	internal_pb "alfred/modules/VCSConnectionService/v1/internal/pb"
	"alfred/modules/VCSConnectionService/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"
	"time"
)

var (
	GithubOAuthScope = []string{"user", "gist", "repo", "repo_deployment", "repo_hook"}
)

func (s *VCSConnectionService) ListAllSupportedVCSProviders(context.Context, *empty.Empty) (*pb.ListAllSupportedVCSProvidersResponse, error) {
	conf := s.config.VCSConfig
	var providers []string
	for _, vcs := range conf {
		providers = append(providers, vcs.Provider)
	}
	return &pb.ListAllSupportedVCSProvidersResponse{
		Providers: providers,
	}, nil
}
func (s *VCSConnectionService) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	var redirectUrl string
	switch in.Provider {
	case pb.VCSConnectionProvider_VCS_GITHUB:
		conf := s.config.VCSConfig
		// Format required field in url
		mp := map[string]string{
			"scope":         conf["vcs_github"].Scope,
			"client_id":     conf["vcs_github"].ClientID,
			"redirect_uri":  conf["vcs_github"].RedirectUri,
			"response_type": conf["vcs_github"].ResponseType,
		}
		reg := regexp.MustCompile(`[{}]`)
		fields := reg.Split(conf["vcs_github"].URLTemplate, -1)
		for i := 1; i < len(fields); i = i + 2 {
			val, ok := mp[fields[i]]
			if ok {
				fields[i] = val
			} else {
				return nil, status.Error(codes.Internal, "Internal ERROR")
			}
		}
		redirectUrl = strings.Join(fields, "")
	}

	// todo - find account id to identify user account
	return &pb.AuthorizeResponse{
		RedirectUrl:  redirectUrl,
		TempJwtToken: "",
	}, nil
}
func (s *VCSConnectionService) Callback(ctx context.Context, in *pb.CallbackRequest) (*pb.AccountVCSConnection, error) {
	// Getting required tokens and their expiry
	var accessToken, refreshToken string
	var accessTokenExpiry, refreshTokenExpiry time.Time

	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "Connection code not provided")
	}

	switch in.Provider {
	case pb.VCSConnectionProvider_VCS_GITHUB:
		//config
		conf := s.config.VCSConfig
		githubOauthConfig := &oauth2.Config{
			RedirectURL:  conf["vcs_github"].RedirectUri,
			ClientID:     conf["vcs_github"].ClientID,
			ClientSecret: conf["vcs_github"].ClientSecret,
			Scopes:       GithubOAuthScope,
			Endpoint:     github.Endpoint,
		}

		gToken, err := githubOauthConfig.Exchange(ctx, in.Code)
		if err != nil {
			//ctxzap.Extract(ctx).Error("Unable to get token", zap.Error(err))
			return nil, status.Error(codes.Internal, "INTERNAL ERROR")
		}

		accessToken = gToken.AccessToken
		refreshToken = gToken.RefreshToken
		accessTokenExpiry = time.Now().UTC().AddDate(1, 11, 0) //todo - fix time
		refreshTokenExpiry = time.Now().UTC().AddDate(1, 11, 0)
	}

	// time.Time type to Timestamp for proto
	aTEP, err := ptypes.TimestampProto(accessTokenExpiry)
	if err != nil {
		return nil, err
	}
	rTEP, err := ptypes.TimestampProto(refreshTokenExpiry)
	if err != nil {
		return nil, err
	}

	vcs := &internal_pb.VCSConnection{
		Provider:           internal_pb.VCSConnectionProvider_VCS_GITHUB,
		ConnectionId:       "",
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  aTEP,
		RefreshTokenExpiry: rTEP,
		Revoked:            false,
		AccountId:          "",
	}
	_, err = s.internalVCSClient.CreateVCSConnection(ctx, &internal_pb.CreateVCSConnectionRequest{
		VcsConnection: vcs,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AccountVCSConnection{
		AccountId: "",
		Provider:  in.Provider,
		Label:     "",
	}, nil
}
func (s *VCSConnectionService) ListVCSConnection(context.Context, *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	return nil, nil
}

//func (s *VCSConnectionService) ListReposistory(ctx context.Context, in *pb.ListReposistoryRequest) (*pb.ListReposistoryResponse, error) {
//	//check the stored accessToken
//	vcs, err := s.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
//		AccountId: "",
//		Provider:  in.Provider,
//	})
//	if err != nil {
//		return nil, err
//	}
//	var accessToken string
//	if vcs.AccessToken != "" {
//		accessToken = vcs.AccessToken
//	} else {
//		return nil, status.Error(codes.NotFound, "Integration token not found")
//	}
//
//	ts := oauth2.StaticTokenSource(
//		&oauth2.Token{AccessToken: accessToken},
//	)
//	tc := oauth2.NewClient(ctx, ts)
//
//	client := gitHub.NewClient(tc)
//
//	// list all repositories for the authenticated user
//	var reposistories []*pb.Reposistory
//	repos, _, err := client.Repositories.List(ctx, "", nil)
//	if err != nil {
//		return nil, err
//	}
//	for _, repo := range repos {
//		reposistory := &pb.Reposistory{
//			Id:       *repo.ID,
//			NodeID:   *repo.NodeID,
//			Name:     *repo.Name,
//			FullName: *repo.FullName,
//			//Description:   *repo.Description,
//			DefaultBranch: *repo.DefaultBranch,
//			//MasterBranch:  *repo.MasterBranch,
//			//CreatedAt:     repo.CreatedAt,
//			//PushedAt:      *repo.PushedAt,
//			//UpdatedAt:     *repo.UpdatedAt,
//			Clone_URL: *repo.CloneURL,
//			Git_URL:   *repo.GitURL,
//			//Size:          *repo.Size,
//			Private:  *repo.Private,
//			Branches: nil,
//		}
//		reposistories = append(reposistories, reposistory)
//	}
//	return &pb.ListReposistoryResponse{
//		Reposistories: reposistories,
//	}, nil
//}
//func (s *VCSConnectionService) GetReposistory(context.Context, *pb.GetReposistoryRequest) (*pb.Reposistory, error) {
//	return nil, nil
//}
