package rest

import (
	"alfred/config"
	"alfred/logger"
	awsArchitecturePb "alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/pb"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	helloWorldPb "alfred/modules/HelloWorldService/pb"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
//Todo: Remove local host from here
func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, config *config.Config) error {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userProfilePb.RegisterUserProfileServiceHandlerFromEndpoint(ctx, mux, config.Grpc.Host+":"+config.Grpc.Port, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := helloWorldPb.RegisterHelloWorldServiceHandlerFromEndpoint(ctx, mux, config.Grpc.Host+":"+config.Grpc.Port, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := architectureSuggesterPb.RegisterArchitectureSuggesterServiceHandlerFromEndpoint(ctx, mux, config.Grpc.Host+":"+config.Grpc.Port, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := awsArchitecturePb.RegisterAwsServiceHandlerFromEndpoint(ctx, mux, config.Grpc.Host+":"+config.Grpc.Port, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	return nil
}
