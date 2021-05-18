package authentication

import (
	"alfred/config"
	account "alfred/modules/account/v1"
	accountPb "alfred/modules/account/v1/pb"
	"alfred/modules/authentication/v1/pb"
	usrProfile "alfred/modules/user-profile/v1"
	usrProfilePb "alfred/modules/user-profile/v1/pb"
	"context"
	oidc "github.com/coreos/go-oidc"
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
	accountInServer     accountPb.AccountInternalServer
	authenticator       *Authenticator
}

// NewAuthenticationServer Service Implementation
func NewAuthenticationServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,

) pb.AuthenticationsServer {
	userProfileSrv := usrProfile.NewUserProfilesServer(db, config, grpcClientConn)
	userProfileInSrv := usrProfile.NewUserProfileInServer(db, config, grpcClientConn)
	accountInSrv := account.NewAccountsInServer(db, config, grpcClientConn)
	authenticatorCli, _ := NewAuthenticator(config)
	return &authenticationServer{
		db:                  db,
		config:              config,
		grpcClient:          grpcClientConn,
		userProfileServer:   userProfileSrv,
		userProfileInServer: userProfileInSrv,
		authenticator:       authenticatorCli,
		accountInServer:     accountInSrv,
	}
}

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

// NewAuthenticator TODO: Move this a Fx
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
