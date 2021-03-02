package graphql

import (
	"alfred/logger"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"context"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
)


//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGraphqlModules(ctx context.Context , mux *runtime.ServeMux) error {
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userProfilePb.RegisterUserProfileServiceGraphql( mux) ;
		err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}
	return nil
}
