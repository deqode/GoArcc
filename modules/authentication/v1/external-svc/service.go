package external_svc

import (
	"context"
	oidc "github.com/coreos/go-oidc"
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/modules/authentication/v1/pb"
	"github.com/deqode/GoArcc/modules/user-profile/v1/external-svc"
	userProfileInt "github.com/deqode/GoArcc/modules/user-profile/v1/internal-svc"
	usrProfilePb "github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

type authenticationServer struct {
	db                  *gorm.DB
	config              *config.Config
	grpcClient          *grpc.ClientConn
	userProfileServer   usrProfilePb.UserProfilesServer
	userProfileInServer usrProfilePb.UserProfileInternalServer
	authenticator       *Authenticator
}

// NewAuthenticationServer Service Implementation
func NewAuthenticationServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,

) pb.AuthenticationsServer {
	userProfileSrv := external_svc.NewUserProfilesServer(db, config, grpcClientConn)
	userProfileInSrv := userProfileInt.NewUserProfileInServer(db, config, grpcClientConn)
	authenticatorCli, _ := NewAuthenticator(config)
	return &authenticationServer{
		db:                  db,
		config:              config,
		grpcClient:          grpcClientConn,
		userProfileServer:   userProfileSrv,
		userProfileInServer: userProfileInSrv,
		authenticator:       authenticatorCli,
	}
}

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator(config *config.Config) (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://"+config.Auth.Auth0Domain+"/")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     config.Auth.Auth0ClientID,
		ClientSecret: config.Auth.Auth0ClientSecret,
		RedirectURL:  config.Auth.Auth0CallBackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "user"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
