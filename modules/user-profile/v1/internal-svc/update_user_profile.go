package internal_svc

import (
	"context"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userProfileInServer) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	if in.GetUserProfile() == nil {
		return nil, status.Error(codes.FailedPrecondition, "UserProfile to update is not provided")
	}
	if err := in.Validate(); err != nil {
		return nil, err
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
