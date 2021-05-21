package internal_svc

import (
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfileInServer) UpdateUserProfile(ctx context.Context, in *pb.UpdateUserProfileRequest) (*pb.UserProfile, error) {
	err := in.Validate()
	if err != nil {
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
