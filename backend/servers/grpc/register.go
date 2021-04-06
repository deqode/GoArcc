package grpc

import (
	"alfred/config"
	"alfred/modules/ArchitectureSuggesterService"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	v13 "alfred/modules/AuthService/v1"
	authPb "alfred/modules/AuthService/v1/pb"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	v12 "alfred/modules/UserProfileService/v1"
	userProfilepb "alfred/modules/UserProfileService/v1/pb"
	"alfred/modules/VCSConnectionService/v1"
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
	userProfilepb.RegisterUserProfileServiceServer(srv, v12.NewUserProfileService(db, config, grpcClientConnection))
	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService())
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, ArchitectureSuggesterService.NewArchitectureSuggesterService())
	authPb.RegisterUserLoginServiceServer(srv, v13.NewAuthService(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionServiceServer(srv, v1.NewVCSConnectionService(db, config, grpcClientConnection))
}
