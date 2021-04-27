package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"strings"
	"time"
)

type claimsKey int

const (
	claimsKeyValue claimsKey = 0
)

func ClaimsFromContext(ctx context.Context) map[string]interface{} {
	v := ctx.Value(claimsKeyValue)
	if v == nil {
		return map[string]interface{}{}
	}

	return v.(map[string]interface{})
}

func WithClaims(ctx context.Context, c map[string]interface{}) context.Context {
	return context.WithValue(ctx, claimsKeyValue, c)
}

func NewJWTUnaryInterceptor(ks *jose.JSONWebKeySet) grpc.UnaryServerInterceptor {
	getToken := func(ctx context.Context) string {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ""
		}

		token := md.Get("authorization")
		if len(token) < 1 {
			return ""
		}

		splitToken := strings.Split(token[0], " ")
		if len(splitToken) < 2 {
			return ""
		}

		return splitToken[1]
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		t := getToken(ctx)
		claims, err := VerifyToken(t, "", []string{}, ks)
		if err != nil {
			return handler(ctx, req)
		}
		//ui := userinfo.FromClaims(claims)
		//if ui.Id != "" {
		//	ctx = metadata.AppendToOutgoingContext(ctx, id, ui.Id, email, ui.Email)
		//}

		return handler(WithClaims(ctx, claims), req)
	}
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

func TokenFromHeader(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) > 1 {
		reqToken = splitToken[1]
		reqToken = strings.TrimSpace(reqToken)
		return reqToken
	}

	return ""
}
