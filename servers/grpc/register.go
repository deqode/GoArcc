package grpc

import (
	"github.com/deqode/GoArcc/config"
	authExt "github.com/deqode/GoArcc/modules/authentication/v1/external-svc"
	authPb "github.com/deqode/GoArcc/modules/authentication/v1/pb"
	userExt "github.com/deqode/GoArcc/modules/user-profile/v1/external-svc"
	userProfilePb "github.com/deqode/GoArcc/modules/user-profile/v1/pb"
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
