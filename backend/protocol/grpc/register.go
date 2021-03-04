package grpc

import (
	"alfred/modules/UserProfileService"
	userProfilepb "alfred/modules/UserProfileService/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Todo : Whenever any new modules will be in alfred : it must be registered in below method
func RegisterGrpcModules(srv *grpc.Server, db *gorm.DB) {
	//register user modules
	userProfilepb.RegisterUserProfileServiceServer(srv, UserProfileService.NewUserProfileService(db))
}
