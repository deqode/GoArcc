package grpc

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/account/v1/external-svc"
	"alfred/modules/account/v1/external-svc/pb"
	"alfred/modules/architecture-suggester"
	architectureSuggesterPb "alfred/modules/architecture-suggester/pb"
	auth "alfred/modules/authentication/v1"
	authPb "alfred/modules/authentication/v1/pb"
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
	userProfilePb.RegisterUserProfilesServer(srv, userProfile.NewUserProfilesServer(db, config, grpcClientConnection))
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, architecture_suggester.NewArchitectureSuggesterService(cadenceConfig, cadenceAdapter))
	authPb.RegisterAuthenticationsServer(srv, auth.NewAuthenticationServer(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionsServer(srv, vcs.NewVCSConnectionServer(db, config, grpcClientConnection))
	pb.RegisterAccountsServer(srv, external_svc.NewAccountExtServer(db, config, grpcClientConnection))
	gitPb.RegisterGitsServer(srv, git.NewGitServer(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
}
