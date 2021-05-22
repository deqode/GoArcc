package external_svc

import (
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"github.com/golang/protobuf/ptypes"
	githubClient "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	GithubOAuthScope = []string{"user", "gist", "repo", "repo_deployment", "repo_hook"}
)

func (s *vcsConnectionServer) Callback(ctx context.Context, in *pb.CallbackRequest) (*pb.AccountVCSConnection, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	// Getting required tokens and their expiry
	var accessToken, refreshToken string
	var accessTokenExpiry, refreshTokenExpiry time.Time
	var user *githubClient.User

	vcs, err := s.vcsConnectionIntServer.GetVCSConnection(ctx, &pb.GetVCSConnectionRequest{
		AccountId: in.AccountId,
		Provider:  in.Provider,
	})
	if vcs != nil {
		return &pb.AccountVCSConnection{
			AccountId: vcs.AccountId,
			Label:     vcs.Label,
			Provider:  vcs.Provider,
		}, nil
	}
	switch in.Provider {
	case types.VCSProviders_GITHUB:
		githubConfig := s.config.GithubVCSConfig
		githubOauthConfig := &oauth2.Config{
			RedirectURL:  githubConfig.RedirectURI,
			ClientID:     githubConfig.ClientID,
			ClientSecret: githubConfig.ClientSecret,
			Scopes:       GithubOAuthScope,
			Endpoint:     github.Endpoint,
		}

		gToken, err := githubOauthConfig.Exchange(ctx, in.Code)
		if err != nil {
			//ctxzap.Extract(ctxhelper).Error("Unable to get token", zap.Error(err))
			return nil, status.Error(codes.Internal, "INTERNAL ERROR")
		}
		oauthClient := githubOauthConfig.Client(oauth2.NoContext, gToken)
		client := githubClient.NewClient(oauthClient)
		user, _, err = client.Users.Get(ctx, "")
		if err != nil {
			return nil, err
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

	label := in.Provider.String() + "-" + *user.Name
	vcsObj := &pb.VCSConnection{
		Provider:           in.Provider,
		ConnectionId:       *user.Login,
		AccessToken:        accessToken,
		RefreshToken:       refreshToken,
		AccessTokenExpiry:  aTEP,
		RefreshTokenExpiry: rTEP,
		Revoked:            false,
		AccountId:          in.AccountId,
		UserName:           *user.Login,
		Label:              label,
	}
	_, err = s.vcsConnectionIntServer.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
		VcsConnection: vcsObj,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AccountVCSConnection{
		AccountId: in.AccountId,
		Provider:  in.Provider,
		Label:     label,
	}, nil
}
