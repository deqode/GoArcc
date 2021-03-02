package implementation

import (
	"alfred/modules/UserProfileService/pb"
	"context"
)

func (u *UserProfileService) GetUserProfile(ctx context.Context, request *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}
