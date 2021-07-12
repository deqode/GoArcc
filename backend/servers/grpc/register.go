package grpc

import (
	"goarcc/config"
	authExt "goarcc/modules/authentication/v1/external-svc"
	authPb "goarcc/modules/authentication/v1/pb"
	userExt "goarcc/modules/user-profile/v1/external-svc"
	userProfilePb "goarcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// RegisterGrpcModules Todo : Whenever any new modules will be in goarcc : it must be registered in below method
/*
  RegisterGrpcModules: will register the modules/services to the server.
*/
func RegisterGrpcModules(srv *grpc.Server,
	db *gorm.DB,
	config *config.Config,
	grpcClientConnection *grpc.ClientConn) {

	//todo register new grpc modules here
	//register user modules
	userProfilePb.RegisterUserProfilesServer(srv, userExt.NewUserProfilesServer(db, config, grpcClientConnection))
	authPb.RegisterAuthenticationsServer(srv, authExt.NewAuthenticationServer(db, config, grpcClientConnection))
}
