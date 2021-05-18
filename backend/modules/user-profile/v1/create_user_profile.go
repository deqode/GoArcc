package user_profile

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfileInServer) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	return s.store.CreateUserProfile(ctx, in)
}
