package rest

import (
	"alfred.sh/common/logger"
	accountPb "alfred/modules/account/v1/pb"
	authPb "alfred/modules/authentication/v1/pb"
	gitPb "alfred/modules/git/v1/pb"
	stackPb "alfred/modules/stack/v1/pb"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterRESTModules Todo : Whenever any new modules will be in alfred : it must be registered in below method
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

	if err := vcsPb.RegisterVCSConnectionsHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := accountPb.RegisterAccountsHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := gitPb.RegisterGitsHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := stackPb.RegisterStacksHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
