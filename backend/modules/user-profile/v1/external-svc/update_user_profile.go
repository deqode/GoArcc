package external_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfileInServer) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	return s.store.UpdateUserProfile(ctx, in)
}
