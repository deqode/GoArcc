package grpc

import (
	grpcClient "alfred/client/grpcClient"
	"alfred/config"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB, config *config.Config) {
	//making client : Note it will run under the same context
	clientContext := grpcClient.GetGrpcClientConnectionContext()
	conn := grpcClient.GetGrpcClientConnection(config, clientContext)

	//todo register new grpc modules here
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db, hellopb.NewHelloWorldServiceClient(conn)))
	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService(userProfilepb.NewUserProfileServiceClient(conn)))
}
