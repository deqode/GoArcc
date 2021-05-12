package internals

import (
	"alfred/modules/user-profile/v1/internals/pb"
	"context"
)

func (s *userProfileInternalServer) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	panic("implement me")
}
