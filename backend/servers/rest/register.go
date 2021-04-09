package rest

import (
	"alfred/logger"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	authPb "alfred/modules/AuthService/v1/pb"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	vcsPb "alfred/modules/VCSConnectionService/v1/pb"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
//Todo: Remove local host from here
func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userProfilePb.RegisterUserProfileServiceHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := architectureSuggesterPb.RegisterArchitectureSuggesterServiceHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := authPb.RegisterUserLoginServiceHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := vcsPb.RegisterVCSConnectionServiceHandler(ctx, mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	return nil
}
