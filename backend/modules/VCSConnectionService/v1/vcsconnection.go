package v1

import (
	"alfred/logger"
	"alfred/modules/VCSConnectionService/v1/models"
	"alfred/modules/VCSConnectionService/v1/pb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	gitHub "github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	GithubOAuthScope = []string{"user", "gist", "repo", "repo_deployment", "repo_hook"}
)

func (s *VCSConnectionService) ListAlfredVCSConnection(context.Context, *empty.Empty) (*pb.ListAlfredVCSConnectionResponse, error) {
	return nil, nil
}
func (s *VCSConnectionService) Connect(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	// finding particular info
	var info *Info
	for _, i := range s.info {
		if i == nil {
			continue
		}
		if i.Provider == in.Provider.String() {
			info = i
			break
		}
	}
	if info == nil {
		return nil, status.Error(codes.InvalidArgument, "Incorrect VCS type")
	}

	// Format required field in url
	mp := map[string]string{
		"scope":         info.Scope,
		"client_id":     info.ClientID,
		"redirect_uri":  info.RedirectUri,
		"response_type": info.ResponseType,
	}
	reg := regexp.MustCompile(`[{}]`)
	fields := reg.Split(info.URLTemplate, -1)
	for i := 1; i < len(fields); i = i + 2 {
		val, ok := mp[fields[i]]
		if ok {
			fields[i] = val
		} else {
			return nil, status.Error(codes.Internal, "Internal ERROR")
		}
	}

	// todo - find account id to identify user account
	return &pb.ConnectResponse{
		Uri:          strings.Join(fields, ""),
		UserId:       "",
		TempJwtToken: "",
		AccountId:    "",
	}, nil
}
func (s *VCSConnectionService) Connected(ctx context.Context, in *pb.ConnectedRequest) (*pb.AccountVCSConnection, error) {
	// Getting required tokens and their expiry
	var accessToken, refreshToken string
	var accessTokenExpiry, refreshTokenExpiry time.Time

	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "Connection code not provided")
	}

	//var info *Info
	//for _, i := range s.info {
	//	if i == nil {
	//		continue
	//	}
	//	if i.Provider == in.Provider.String() {
	//		info = i
	//		break
	//	}
	//}

	if info == nil {
		return nil, status.Error(codes.InvalidArgument, "Incorrect VCS type")
	}

	switch i := in.Provider; i {
	case pb.VCSConnectionProvider_VCS_GITHUB:
		// config
		githubOauthConfig := &oauth2.Config{
			RedirectURL:  info.RedirectUri,
			ClientID:     info.ClientID,
			ClientSecret: info.ClientSecret,
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

	vcs := &pb.VCSConnection{
		Provider:           pb.VCSConnectionProvider_VCS_GITHUB,
		ConnectionId:       "",
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  aTEP,
		RefreshTokenExpiry: rTEP,
		Revoked:            false,
		AccountId:          "",
	}
	_, err = s.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
		VcsConnection: vcs,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AccountVCSConnection{
		AccountId: "",
		Provider:  in.Provider,
		Id:        "",
	}, nil
}
func (s *VCSConnectionService) ListVCSConnection(context.Context, *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {
	return nil, nil
}
func (s *VCSConnectionService) RevokeVCSToken(context.Context, *pb.RevokeVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *VCSConnectionService) RenewVCSToken(context.Context, *pb.RenewVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *VCSConnectionService) ListReposistory(ctx context.Context, in *pb.ListReposistoryRequest) (*pb.ListReposistoryResponse, error) {
	//check the stored accessToken
	vcs, err := s.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
		AccountId: "",
		Provider:  in.Provider,
	})
	if err != nil {
		return nil, err
	}
	var accessToken string
	if vcs.AccessToken != "" {
		accessToken = vcs.AccessToken
	} else {
		return nil, status.Error(codes.NotFound, "Integration token not found")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := gitHub.NewClient(tc)

	// list all repositories for the authenticated user
	var reposistories []*pb.Reposistory
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		reposistory := &pb.Reposistory{
			Id:       *repo.ID,
			NodeID:   *repo.NodeID,
			Name:     *repo.Name,
			FullName: *repo.FullName,
			//Description:   *repo.Description,
			DefaultBranch: *repo.DefaultBranch,
			//MasterBranch:  *repo.MasterBranch,
			//CreatedAt:     repo.CreatedAt,
			//PushedAt:      *repo.PushedAt,
			//UpdatedAt:     *repo.UpdatedAt,
			Clone_URL: *repo.CloneURL,
			Git_URL:   *repo.GitURL,
			//Size:          *repo.Size,
			Private:  *repo.Private,
			Branches: nil,
		}
		reposistories = append(reposistories, reposistory)
	}
	return &pb.ListReposistoryResponse{
		Reposistories: reposistories,
	}, nil
}
func (s *VCSConnectionService) GetReposistory(context.Context, *pb.GetReposistoryRequest) (*pb.Reposistory, error) {
	return nil, nil
}
func (s *VCSConnectionService) CreateVCSConnection(ctx context.Context, in *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error) {
	//creating uuid
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Log.Debug("unable to generate uuid")
	}
	aTEP, err := ptypes.Timestamp(in.VcsConnection.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}
	rTEP, err := ptypes.Timestamp(in.VcsConnection.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}
	VCSModel := &models.VCSConnection{
		ID:                 fmt.Sprintf("%s", out),
		Provider:           1,
		ConnectionId:       in.VcsConnection.ConnectionId,
		AccessToken:        in.VcsConnection.AccessToken,
		RefreshToken:       in.VcsConnection.RefreshToken,
		AccessTokenExpiry:  &aTEP,
		RefreshTokenExpiry: &rTEP,
		Revoked:            false,
		AccountId:          in.VcsConnection.AccountId,
	}

	t := s.db.Create(VCSModel)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.VcsConnection, nil
}


