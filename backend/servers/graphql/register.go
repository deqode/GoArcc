package graphql

import (
	"alfred/logger"
	awsArchitecturePb "alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/pb"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	hellopb "alfred/modules/HelloWorldService/pb"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGraphqlModules: Mapping the services with the single graphql endpoint
*/
func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := userProfilePb.RegisterUserProfileServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := hellopb.RegisterHelloWorldServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := architectureSuggesterPb.RegisterArchitectureSuggesterServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := awsArchitecturePb.RegisterAwsServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
