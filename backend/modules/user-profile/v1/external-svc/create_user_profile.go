package external_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userProfileInServer) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if in.GetUserProfile() == nil {
		return nil, status.Error(codes.FailedPrecondition, "UserProfile not provided")
	}

	return s.store.CreateUserProfile(ctx, in)
}
