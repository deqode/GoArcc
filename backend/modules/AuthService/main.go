package AuthService

import (
	"alfred/config"
	"alfred/modules/AuthService/pb"
	"context"
	"golang.org/x/oauth2"
	"log"

	oidc "github.com/coreos/go-oidc"
)

type AuthService struct {
}

//Service Implementation
func NewAuthService() pb.UserLoginServiceServer {
	return &AuthService{}
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
		ClientID:     config.Auth.Auth0ClientId,
		ClientSecret: config.Auth.Auth0ClientSecret,
		RedirectURL:  config.Auth.Auth0CallbackUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
