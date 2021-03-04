package grpc

import (
	"alfred/config"
	"alfred/logger"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB, config *config.Config) {
	//making client : Note it will run under the same context
	conn, err := grpc.Dial("localhost:"+config.GRPCPort, grpc.WithInsecure())
	if err != nil {
		logger.Log.Fatal("did not connect", zap.Error(err))
	}

	//todo register new grpc modules here
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db, hellopb.NewHelloWorldServiceClient(conn)))
	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService(userProfilepb.NewUserProfileServiceClient(conn)))
}
