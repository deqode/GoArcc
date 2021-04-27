package UserProfileService

import (
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserProfileService) GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error) {
	profile := models.UserProfile{}
	result := s.db.First(&profile, "sub = ?", in.Sub)

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	return &pb.UserProfile{
		Id:             profile.ID,
		ExternalSource: profile.Source,
		Name:           profile.Name,
		Email:          profile.Email,
		Sub:            profile.Sub,
		UserName:       profile.UserName,
	}, nil
}
