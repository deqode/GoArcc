package UserProfileService

import (
	"alfred/modules/UserProfileService/v1/models"
	"alfred/modules/UserProfileService/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserProfileService) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	profile := models.UserProfile{}
	result := s.db.First(&profile, "id = ?", in.Id)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return &pb.UserProfile{
		Id:             profile.ID,
		ExternalSource: profile.Source,
		Name:           profile.Name,
		Email:          profile.Email,
	}, nil
}