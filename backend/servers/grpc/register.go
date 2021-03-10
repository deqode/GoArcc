package grpc

import (
	"alfred/config"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB, config *config.Config, conn *grpc.ClientConn) {
	//todo register new grpc modules here
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db, hellopb.NewHelloWorldServiceClient(conn)))

	//hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService(userProfilepb.NewUserProfileServiceClient(conn)))
}
