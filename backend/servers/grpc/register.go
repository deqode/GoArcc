package grpc

import (
	"alfred/config"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGrpcModules: will register the modules/services to the server.
*/
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB, config *config.Config) {
	//todo register new grpc modules here
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db, config))

	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService())
}
