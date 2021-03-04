package graphql

import (
	"alfred/logger"
	hellopb "alfred/modules/HelloWorldService/pb"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"context"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGraphqlModules(ctx context.Context, mux *runtime.ServeMux) error {
	if err := userProfilePb.RegisterUserProfileServiceGraphql(mux); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := hellopb.RegisterHelloWorldServiceGraphql(mux); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
