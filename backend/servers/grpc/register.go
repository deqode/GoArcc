package grpc

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/account/v1/external-svc"
	"alfred/modules/account/v1/external-svc/pb"
	"alfred/modules/architecture-suggester"
	architectureSuggesterPb "alfred/modules/architecture-suggester/pb"
	external_svc2 "alfred/modules/authentication/v1/external-svc"
	authPb "alfred/modules/authentication/v1/pb"
	external_svc3 "alfred/modules/git/v1/external-svc"
	gitPb "alfred/modules/git/v1/pb"
	external_svc4 "alfred/modules/user-profile/v1/external-svc"
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
	userProfilePb.RegisterUserProfilesServer(srv, external_svc4.NewUserProfilesServer(db, config, grpcClientConnection))
	architectureSuggesterPb.RegisterArchitectureSuggesterServiceServer(srv, architecture_suggester.NewArchitectureSuggesterService(cadenceConfig, cadenceAdapter))
	authPb.RegisterAuthenticationsServer(srv, external_svc2.NewAuthenticationServer(db, config, grpcClientConnection))
	vcsPb.RegisterVCSConnectionsServer(srv, vcs.NewVCSConnectionServer(db, config, grpcClientConnection))
	pb.RegisterAccountsServer(srv, external_svc.NewAccountExtServer(db, config, grpcClientConnection))
	gitPb.RegisterGitsServer(srv, external_svc3.NewGitServer(db, config, grpcClientConnection, cadenceConfig, cadenceAdapter))
}
