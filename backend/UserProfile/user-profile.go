package UserProfile

import (
	"alfred/UserProfile/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserServer struct {

}

func NewUserServer() pb.UserProfileServiceServer {
	return &UserServer{
	}
}

func (u *UserServer) CreateUserProfile(ctx context.Context, request *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {

	return &pb.UserProfile{
		Name:   "Atul",
		Id:     "jkwdjnksndkd",
		Source: pb.SOURCE_BITBUCKET,
	}, nil
}

func (u *UserServer) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}

func (u *UserServer) GetUserProfile(ctx context.Context, request *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}

func (u *UserServer) DeleteUserProfile(ctx context.Context, request *pb.DeleteUserProfileRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (u *UserServer) mustEmbedUnimplementedUserProfileServiceServer() {
	panic("implement me")
}
