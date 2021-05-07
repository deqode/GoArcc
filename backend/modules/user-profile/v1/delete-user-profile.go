package UserProfileService

import (
	"alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *UserProfileService) DeleteUserProfile(ctx context.Context, request *pb.DeleteUserProfileRequest) (*empty.Empty, error) {

	// this will be soft delete from the system
	t := s.db.Delete(&models.UserProfile{
		ID: request.Id,
	})
	if t.Error != nil {
		return nil, t.Error
	}
	return nil, nil
}
