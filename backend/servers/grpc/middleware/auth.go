package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"strings"
	"time"
)

//
func myAuthFunction(ctx context.Context) (context.Context, error) {
	authRequired := true
	//getting auth from the context . ie bearer
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "authorization failed")
	}
	//get method ame
	operationName, _ := grpc.Method(ctx)
	//check for public end point
	for _, v := range public_endpoint {
		if operationName == v {
			authRequired = false
		}
	}
	if authRequired {
		token := md.Get("authorization")
		if len(token) < 1 {
			return nil, status.Error(codes.FailedPrecondition, "authorization failed")
		}
		splitToken := strings.Split(token[0], " ")
		if len(splitToken) < 2 {
			return nil, status.Error(codes.FailedPrecondition, "authorization failed")
		}
		//bearerToken := splitToken[1]
	}

	return ctx, nil
}

func VerifyToken(t, iss string, aud []string, ks *jose.JSONWebKeySet) (map[string]interface{}, error) {
	token, err := jwt.ParseSigned(t)
	if err != nil {
		return nil, err
	}

	claims := jwt.Claims{}
	cmap := make(map[string]interface{})
	token.Claims(ks, &claims, &cmap)
	if err := claims.Validate(jwt.Expected{
		Issuer:   iss,
		Time:     time.Now(),
		Audience: aud,
	}); err != nil {
		return nil, err
	}

	return cmap, nil
}
