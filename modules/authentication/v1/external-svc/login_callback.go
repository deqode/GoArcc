package external_svc

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/deqode/GoArcc/modules/authentication/v1/pb"
	userprofileModel "github.com/deqode/GoArcc/modules/user-profile/v1/models"
	userProfilePb "github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"github.com/deqode/GoArcc/protos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

const tokenKey = "id_token"

func (s *authenticationServer) LoginCallback(ctx context.Context, in *pb.LoginCallbackRequest) (*pb.LoginCallbackResponse, error) {
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request is Nil")
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}
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
	var userId string
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
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	userId, err = s.CreateUserAndAccount(ctx, profile)
	if err != nil {
		return nil, err
	}
	return &pb.LoginCallbackResponse{
		IdToken:     rawIDToken,
		AccessToken: token.AccessToken,
		UserId:      userId,
	}, nil
}

func (s *authenticationServer) CreateUserAndAccount(ctx context.Context, profile map[string]interface{}) (string, error) {
	gormDb := s.db
	userid := ""
	err := gormDb.Transaction(func(transaction *gorm.DB) error {
		usr := struct {
			ID            string
			Name          string
			UserName      string
			Sub           string
			ProfilePicURL string
			Source        types.VCSProviders
		}{
			ID:            fmt.Sprintf("%s", profile["sub"]),
			Name:          fmt.Sprintf("%s", profile["name"]),
			UserName:      fmt.Sprintf("%s", profile["nickname"]),
			Sub:           fmt.Sprintf("%s", profile["sub"]),
			ProfilePicURL: fmt.Sprintf("%s", profile["picture"]),
			Source:        types.VCSProviders_GITHUB,
		}
		userProfileModel := userprofileModel.UserProfile{
			ID:            usr.ID,
			Name:          usr.Name,
			UserName:      usr.UserName,
			Sub:           usr.Sub,
			ProfilePicURL: usr.ProfilePicURL,
			Source:        usr.Source,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     gorm.DeletedAt{},
		}
		tx := transaction.Create(&userProfileModel)
		if tx.Error != nil {
			return tx.Error
		}

		//if everything goes right then the transaction will commit by itself
		userid = usr.ID
		return nil
	})
	//when transaction does not happen for some internal reason
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	return userid, nil
}
