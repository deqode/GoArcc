package internal_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userProfileInServer) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	if in.GetUserProfile() == nil {
		return nil, status.Error(codes.FailedPrecondition, "UserProfile to update is not provided")
	}
	return &pb.UserProfile{
		Id:             "",
		Sub:            "",
		Name:           "",
		UserName:       "",
		Email:          "",
		PhoneNumber:    "",
		ExternalSource: 0,
		ProfilePicUrl:  "",
		TokenValidTill: nil,
	}, nil
}
