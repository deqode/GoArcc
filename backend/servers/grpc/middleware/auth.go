package middleware

import (
	"context"
	"encoding/json"
	"github.com/justinas/alice"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"strings"
	"time"

	"alfred/util/userinfo"
)

const (
	id                       = "id"
	email                    = "email"
	claimsKeyValue claimsKey = 0
)

type (
	claimsKey int
)

type Config struct {
	Issuer   string
	Secret   string
	Domain   string
	Audience []string
}

// NewConfig todo : keep this config into config-local
// todo: Move these base AUth0 URLs to config
func NewConfig() (Config, error) {
	newConfig := Config{
		Issuer:   "https://alfred-sh.us.auth0.com/",
		Secret:   "secret",
		Domain:   "https://alfred-sh.us.auth0.com",
		Audience: []string{"BFnfdaibKSdqkSAOksr3XuUNJuCW9zbZ"},
	}

	return newConfig, nil
}

//TODO - User Related Validation
func AuthMiddleware(ctx context.Context) (context.Context, error) {
	authRequired := true
	//getting authentication from the context . ie bearer
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "authorization failed")
	}
	//get method name
	operationName, _ := grpc.Method(ctx)
	//check for public end point
	for _, v := range publicEndpoint {
		if operationName == v {
			authRequired = false
			break
		}
	}
	if authRequired {
		token := md.Get("authorization")
		if len(token) < 1 {
			return nil, status.Error(codes.FailedPrecondition, "Authorization header missing")
		}
		splitToken := strings.Split(token[0], " ")
		if len(splitToken) < 2 {
			return nil, status.Error(codes.FailedPrecondition, "Bearer token not found.")
		}
		jwk, err := NewJwk()
		if err != nil {
			return nil, err
		}
		_, err = NewJWtMiddleware(jwk.JSONWebKeySet, splitToken[1])
		if err != nil {
			return nil, err
		}
		ctx, err = NewJWTUnaryInterceptor(ctx, jwk.JSONWebKeySet, splitToken[1])
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

type DiscoveryResponse struct {
	fx.Out
	RevocationEndpoint string `name:"revocation_endpoint"`
	JSONWebKeySet      *jose.JSONWebKeySet
}

// NewJwk NewJwk() get certificate from openIdConnect
func NewJwk() (DiscoveryResponse, error) {
	c, _ := NewConfig()
	discoveryURI := strings.TrimSuffix(c.Issuer, "/") + "/.well-known/openid-configuration"

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(discoveryURI)
	if err != nil {
		return DiscoveryResponse{}, err
	}
	defer res.Body.Close()

	var discoResp struct {
		JwkURI             string `json:"jwks_uri"`
		RevocationEndpoint string `json:"revocation_endpoint"`
	}

	if err := json.NewDecoder(res.Body).Decode(&discoResp); err != nil {
		return DiscoveryResponse{}, err
	}

	res, err = client.Get(discoResp.JwkURI)
	if err != nil {
		return DiscoveryResponse{}, err
	}
	defer res.Body.Close()

	var keySet jose.JSONWebKeySet
	if err := json.NewDecoder(res.Body).Decode(&keySet); err != nil {
		return DiscoveryResponse{}, err
	}

	return DiscoveryResponse{
		RevocationEndpoint: discoResp.RevocationEndpoint,
		JSONWebKeySet:      &keySet,
	}, nil
}

// NewJWtMiddleware verify the user's token
func NewJWtMiddleware(ks *jose.JSONWebKeySet, token string) (alice.Constructor, error) {
	c, _ := NewConfig()
	_, err := VerifyToken(token, c.Issuer, c.Audience, ks)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unable to verify User")
	}
	return nil, nil
}

// NewJWTUnaryInterceptor add users claims to outgoing context
func NewJWTUnaryInterceptor(ctx context.Context, ks *jose.JSONWebKeySet, token string) (context.Context, error) {
	c, _ := NewConfig()
	claims, err := VerifyToken(token, c.Issuer, c.Audience, ks)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unable to verify User")
	}

	ui := userinfo.FromClaims(claims)
	if ui.ID != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, id, ui.ID, email, ui.Email)
	}
	return WithClaims(ctx, claims), nil
}

func VerifyToken(t, iss string, aud []string, ks *jose.JSONWebKeySet) (map[string]interface{}, error) {
	token, err := jwt.ParseSigned(t)
	if err != nil {
		return nil, err
	}

	claims := jwt.Claims{}
	claimMap := make(map[string]interface{})
	err = token.Claims(ks, &claims, &claimMap)
	if err != nil {
		return nil, err
	}
	if err := claims.Validate(jwt.Expected{
		Issuer:   iss,
		Time:     time.Now(),
		Audience: aud,
	}); err != nil {
		return nil, err
	}

	return claimMap, nil
}

func WithClaims(ctx context.Context, c map[string]interface{}) context.Context {
	return context.WithValue(ctx, claimsKeyValue, c)
}
