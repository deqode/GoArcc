package implementation

import (
	"alfred/modules/UserProfile/pb"
	"context"
)

func (server *UserServer) CreateUserProfile(ctx context.Context, request *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {

	return &pb.UserProfile{
		Name:   "Atul",
		Id:     "jkwdjnksndkd",
		Source: pb.SOURCE_BITBUCKET,
	}, nil
}
