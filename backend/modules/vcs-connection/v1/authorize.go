package vcs_connection

import (
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"
	"time"
)

func (s *vcsConnectionServer) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {

	var redirectURL string
	switch in.Provider {
	case types.VCSProviders_GITHUB:
		githubConfig := s.config.GithubVCSConfig
		// Format required field in url
		mp := map[string]string{
			"scope":         githubConfig.Scope,
			"client_id":     githubConfig.ClientID,
			"redirect_uri":  githubConfig.RedirectURI,
			"response_type": githubConfig.ResponseType,
		}
		reg := regexp.MustCompile(`[{}]`)
		fields := reg.Split(githubConfig.URLTemplate, -1)
		for i := 1; i < len(fields); i = i + 2 {
			val, ok := mp[fields[i]]
			if ok {
				fields[i] = val
			} else {
				return nil, status.Error(codes.Internal, "Internal ERROR")
			}
		}
		redirectURL = strings.Join(fields, "")
	}

	// todo - find account id to identify user account
	return &pb.AuthorizeResponse{
		RedirectUrl:  redirectURL,
		TempJwtToken: "",
	}, nil
}

type UserClaims struct {
	jwt.StandardClaims
	UserID    string `json:"userid"`
	AccountID string `json:"accountid"`
}

//TODO- use this to send temporary jwt token
// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func Generate(userid, accountid string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		AccountID: accountid,
		UserID:    userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
