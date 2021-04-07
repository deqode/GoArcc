package UserProfileService

import (
	"alfred/modules/UserProfileService/v1/pb"
	"context"
)

func (u *UserProfileService) GetUserProfile(ctx context.Context, request *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}
