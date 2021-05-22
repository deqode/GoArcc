package grpc

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/account/v1/external-svc"
	accountPb "alfred/modules/account/v1/pb"
	authExt "alfred/modules/authentication/v1/external-svc"
	authPb "alfred/modules/authentication/v1/pb"
	gitExt "alfred/modules/git/v1/external-svc"
	gitPb "alfred/modules/git/v1/pb"
	userExt "alfred/modules/user-profile/v1/external-svc"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	vcsExt "alfred/modules/vcs-connection/v1/external-svc"
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
	userProfilePb.RegisterUserProfilesServer(srv, userExt.NewUserProfilesServer(db, config, grpcClientConnection))
	authPb.RegisterAuthenticationsServer(srv, authExt.NewAuthenticationServer(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionsServer(srv, vcsExt.NewVCSConnectionServer(db, config, grpcClientConnection))
	accountPb.RegisterAccountsServer(srv, external_svc.NewAccountExtServer(db, config, grpcClientConnection))
	gitPb.RegisterGitsServer(srv, gitExt.NewGitServer(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
}
