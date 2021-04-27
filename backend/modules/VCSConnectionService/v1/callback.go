package VCSConnectionService

import (
	internal_pb "alfred/modules/VCSConnectionService/v1/internals/pb"
	"alfred/modules/VCSConnectionService/v1/pb"
	"alfred/protos/types"
	"context"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *VCSConnectionService) Callback(ctx context.Context, in *pb.CallbackRequest) (*pb.AccountVCSConnection, error) {
	// Getting required tokens and their expiry
	var accessToken, refreshToken string
	var accessTokenExpiry, refreshTokenExpiry time.Time

	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "Connection code not provided")
	}

	switch in.Provider {
	case types.GitProviders_GITHUB:
		//config
		conf := s.config.VCSConfig
		githubOauthConfig := &oauth2.Config{
			RedirectURL:  conf["github"].RedirectUri,
			ClientID:     conf["github"].ClientID,
			ClientSecret: conf["github"].ClientSecret,
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
		Provider:           types.GitProviders_GITHUB,
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
