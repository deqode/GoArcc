package grpc

import (
	"alfred/config"
	"alfred/modules/ArchitectureSuggesterService"
	"alfred/modules/ArchitectureSuggesterService/AwsArchitectureService"
	awsArchitecturePb "alfred/modules/ArchitectureSuggesterService/AwsArchitectureService/pb"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	"alfred/modules/AuthService"
	authPb "alfred/modules/AuthService/pb"
	"alfred/modules/HelloWorldService"
	hellopb "alfred/modules/HelloWorldService/pb"
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"alfred/modules/VCSConnectionService"
	vcsPb "alfred/modules/VCSConnectionService/pb"
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
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db, config, grpcClientConnection))
	hellopb.RegisterHelloWorldServiceServer(srv, HelloWorldService.NewHelloService())
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, ArchitectureSuggesterService.NewArchitectureSuggesterService())
	awsArchitecturePb.RegisterAwsServiceServer(srv, AwsArchitectureService.NewAwsArchitectureService())
	authPb.RegisterUserLoginServiceServer(srv, AuthService.NewAuthService(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionServiceServer(srv, VCSConnectionService.NewVCSConnectionService(db, config, grpcClientConnection))
}
