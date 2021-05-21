package external_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *userProfileInServer) DeleteUserProfile(ctx context.Context, in *pb.DeleteUserProfileRequest) (*empty.Empty, error) {

	err := in.Validate()
	if err != nil {
		return nil, err
	}

	return s.store.DeleteUserProfile(ctx, in)
}
