package implementation

import (
	"alfred/modules/UserProfileService/pb"
	"context"
)

func (server *UserProfileService) CreateUserProfile(ctx context.Context, request *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {

	return &pb.UserProfile{
		Name:   "Atul",
		Id:     "jkwdjnksndkd",
		Source: pb.SOURCE_BITBUCKET,
	}, nil
}
