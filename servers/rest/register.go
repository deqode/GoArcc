package rest

import (
	"context"
	"github.com/deqode/GoArcc/logger"
	authPb "github.com/deqode/GoArcc/modules/authentication/v1/pb"
	userProfilePb "github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterRESTModules Todo : Whenever any new modules will be in goarcc : it must be registered in below method
//Todo: Remove local host from here
func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userProfilePb.RegisterUserProfilesHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := authPb.RegisterAuthenticationsHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
