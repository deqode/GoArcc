package rest

import (
	"alfred/logger"
	helloWorldPb "alfred/modules/HelloWorldService/pb"
	userProfilePb "alfred/modules/UserProfileService/pb"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, grpcPort string) error {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := userProfilePb.RegisterUserProfileServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := helloWorldPb.RegisterHelloWorldServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
