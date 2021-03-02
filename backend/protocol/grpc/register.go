package grpc

import (
	userProfile "alfred/UserProfile"
	userProfilepb "alfred/UserProfile/pb"
	"google.golang.org/grpc"
)


//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(server *grpc.Server) {
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(server , userProfile.NewUserServer())
}
