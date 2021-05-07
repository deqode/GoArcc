package grpc

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	acc "alfred/modules/account/v1"
	accPb "alfred/modules/account/v1/pb"
	"alfred/modules/architecture-suggester"
	architectureSuggesterPb "alfred/modules/architecture-suggester/pb"
	v13 "alfred/modules/auth/v1"
	authPb "alfred/modules/auth/v1/pb"
	git "alfred/modules/git/v1"
	gitPb "alfred/modules/git/v1/pb"
	userProfile "alfred/modules/user-profile/v1"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	vcs "alfred/modules/vcs-connection/v1"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
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
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, architecture_suggester.NewArchitectureSuggesterService(cadenceConfig, cadenceAdapter))
	authPb.RegisterUserLoginServiceServer(srv, v13.NewAuthService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	vcsPb.RegisterVCSConnectionsServer(srv, vcs.NewVCSConnectionServer(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	accPb.RegisterAccountsServer(srv, acc.NewAccountsServer(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
	gitPb.RegisterGitServiceServer(srv, git.NewGitService(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
}
