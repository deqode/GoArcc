package external_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfilesServer) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	return s.store.GetUserProfile(ctx, in)
}
