package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
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
