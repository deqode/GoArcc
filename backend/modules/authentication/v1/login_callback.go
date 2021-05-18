package authentication

import (
	accountInPb "alfred/modules/account/v1/pb"
	"alfred/modules/authentication/v1/pb"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const tokenKey = "id_token"

func (s *authenticationServer) LoginCallback(ctx context.Context, in *pb.LoginCallbackRequest) (*pb.LoginCallbackResponse, error) {
	token, err := s.authenticator.Config.Exchange(ctx, in.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	rawIDToken, ok := token.Extra(tokenKey).(string)
	if !ok {
		return nil, status.Error(codes.Internal, "No id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: s.config.Auth.Auth0ClientID,
	}

	idToken, err := s.authenticator.Provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to verify ID Token")
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	//get user_details
	usr, err := s.userProfileServer.GetUserProfileBySub(ctx, &userProfilePb.GetUserProfileBySubRequest{
		Sub: fmt.Sprintf("%s", profile["sub"]),
	})
	if usr != nil {
		return &pb.LoginCallbackResponse{
			IdToken:     rawIDToken,
			AccessToken: token.AccessToken,
			UserId:      usr.Id,
		}, nil
	}
	if err != nil {
		code, _ := status.FromError(err)
		if code.Code() != codes.NotFound {
			return nil, err
		}
	}
	usr.Id, err = s.CreateUserAndAccount(ctx, profile)
	if err != nil {
		return nil, err
	}
	return &pb.LoginCallbackResponse{
		IdToken:     rawIDToken,
		AccessToken: token.AccessToken,
		UserId:      usr.Id,
	}, nil
}

func (s *authenticationServer) CreateUserAndAccount(ctx context.Context, profile map[string]interface{}) (string, error) {
	//create User Profile
	user, err := s.userProfileInServer.CreateUserProfile(ctx, &userProfilePb.CreateUserProfileRequest{
		UserProfile: &userProfilePb.UserProfile{
			Id:             fmt.Sprintf("%s", profile["sub"]),
			Sub:            fmt.Sprintf("%s", profile["sub"]),
			Name:           fmt.Sprintf("%s", profile["name"]),
			UserName:       fmt.Sprintf("%s", profile["nickname"]),
			ProfilePicUrl:  fmt.Sprintf("%s", profile["picture"]),
			ExternalSource: types.VCSProviders_GITHUB,
			TokenValidTill: nil,
		},
	})
	if err != nil {
		return "", err
	}

	//create User Account
	_, err = s.accountInServer.CreateAccount(ctx, &accountInPb.CreateAccountRequest{
		Account: &accountInPb.Account{
			Slug:   user.Name + "_" + user.ExternalSource.String(),
			UserId: user.Id,
		},
	})
	if err != nil {
		return "", err
	}
	return user.GetId(), nil
}
