package grpc

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	acc "alfred/modules/AccountService/v1"
	accPb "alfred/modules/AccountService/v1/pb"
	"alfred/modules/ArchitectureSuggesterService"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	v13 "alfred/modules/AuthService/v1"
	authPb "alfred/modules/AuthService/v1/pb"
	git "alfred/modules/GitService/v1"
	gitPb "alfred/modules/GitService/v1/pb"
	userProfile "alfred/modules/UserProfileService/v1"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	vcs "alfred/modules/VCSConnectionService/v1"
	vcsPb "alfred/modules/VCSConnectionService/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// RegisterGrpcModules Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGrpcModules: will register the modules/services to the server.
*/
func RegisterGrpcModules(srv *grpc.Server,
	db *gorm.DB,
	config *config.Config,
	grpcClientConnection *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter) {
	//todo register new grpc modules here
	//register user modules
	userProfilePb.RegisterUserProfileServiceServer(srv, userProfile.NewUserProfileService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, ArchitectureSuggesterService.NewArchitectureSuggesterService(cadenceConfig, cadenceAdapter))
	authPb.RegisterUserLoginServiceServer(srv, v13.NewAuthService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	vcsPb.RegisterVCSConnectionServiceServer(srv, vcs.NewVCSConnectionService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	accPb.RegisterAccountServiceServer(srv, acc.NewAccountService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	gitPb.RegisterGitServiceServer(srv, git.NewGitService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
}
