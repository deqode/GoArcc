package grpc

import (
	"alfred/config"
	"alfred/modules/ArchitectureSuggesterService"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	v13 "alfred/modules/AuthService/v1"
	authPb "alfred/modules/AuthService/v1/pb"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	userProfile "alfred/modules/UserProfileService/v1"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	vcs "alfred/modules/VCSConnectionService/v1"
	vcsPb "alfred/modules/VCSConnectionService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGrpcModules: will register the modules/services to the server.
*/
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB, config *config.Config, grpcClientConnection *grpc.ClientConn) {
	//todo register new grpc modules here
	//register user modules
	userProfilePb.RegisterUserProfileServiceServer(srv, userProfile.NewUserProfileService(db, config, grpcClientConnection))
	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService())
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, ArchitectureSuggesterService.NewArchitectureSuggesterService())
	authPb.RegisterUserLoginServiceServer(srv, v13.NewAuthService(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionServiceServer(srv, vcs.NewVCSConnectionService(db, config, grpcClientConnection))
}
