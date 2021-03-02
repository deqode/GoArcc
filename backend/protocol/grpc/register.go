package grpc

import (
	userProfile "alfred/modules/UserProfile/implementation"
	userProfilepb "alfred/modules/UserProfile/pb"
)




//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules() {
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(server , userProfile.NewUserServer())
}
