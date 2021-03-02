package grpc

import (
	userProfile "alfred/modules/UserProfileService/implementation"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"google.golang.org/grpc"
)




//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(srv *grpc.Server) {
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv , userProfile.NewUserProfileService())
}
