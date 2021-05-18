package user_profile

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfilesServer) GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	request := &pb.GetUserProfileRequest{
		Id: in.Sub,
	}
	return s.store.GetUserProfile(ctx, request)
}
