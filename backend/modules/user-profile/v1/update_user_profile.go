package user_profile

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfileServer) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}
