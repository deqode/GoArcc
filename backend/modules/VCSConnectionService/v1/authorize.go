package VCSConnectionService

import (
	"alfred/modules/VCSConnectionService/v1/pb"
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

func (s *VCSConnectionService) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	var redirectUrl string
	switch in.Provider {
	case types.GitProviders_GITHUB:
		conf := s.config.VCSConfig
		// Format required field in url
		mp := map[string]string{
			"scope":         conf["github"].Scope,
			"client_id":     conf["github"].ClientID,
			"redirect_uri":  conf["github"].RedirectUri,
			"response_type": conf["github"].ResponseType,
		}
		reg := regexp.MustCompile(`[{}]`)
		fields := reg.Split(conf["github"].URLTemplate, -1)
		for i := 1; i < len(fields); i = i + 2 {
			val, ok := mp[fields[i]]
			if ok {
				fields[i] = val
			} else {
				return nil, status.Error(codes.Internal, "Internal ERROR")
			}
		}
		redirectUrl = strings.Join(fields, "")
	}

	// todo - find account id to identify user account
	return &pb.AuthorizeResponse{
		RedirectUrl:  redirectUrl,
		TempJwtToken: "",
	}, nil
}

type UserClaims struct {
	jwt.StandardClaims
	userId    string `json:"userid"`
	accountId string `json:"accountid"`
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

func Generate(userid, accountid string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		accountId: accountid,
		userId:    userid,
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
