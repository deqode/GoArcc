package AuthService

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/AuthService/v1/pb"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	"context"
	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"time"
)

type AuthService struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	userProfileClient userProfilePb.UserProfileServiceClient
	authenticator     *Authenticator
	cadenceConfig     *background.CadenceAppConfig
	cadenceAdapter    *cadenceAdapter.CadenceAdapter
}

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

// NewAuthService Service Implementation
func NewAuthService(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.UserLoginServiceServer {
	userProfileCli := userProfilePb.NewUserProfileServiceClient(grpcClientConn)
	authenticatorCli, _ := NewAuthenticator(config)
	return &AuthService{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		userProfileClient: userProfileCli,
		authenticator:     authenticatorCli,
		cadenceAdapter:    cadenceAdapter,
		cadenceConfig:     cadenceConfig,
	}
}

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

// TODO: Move this a Fx
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
		RedirectURL:  config.Auth.Auth0CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "user"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
