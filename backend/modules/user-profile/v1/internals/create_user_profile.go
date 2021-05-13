package internals

import (
	"alfred/modules/user-profile/v1/internals/pb"
	"alfred/modules/user-profile/v1/models"
	"context"
)

func (s *userProfileInternalServer) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	//id, err := uuid.GenerateUUID()
	//if err != nil {
	//	return nil, err
	//}

	usrProfile := &models.UserProfile{
		ID:            in.UserProfile.Sub,
		Name:          in.UserProfile.Name,
		Email:         in.UserProfile.Email,
		PhoneNumber:   in.UserProfile.PhoneNumber,
		Sub:           in.UserProfile.Sub,
		UserName:      in.UserProfile.UserName,
		ProfilePicURL: in.UserProfile.ProfilePicUrl,
		Source:        in.UserProfile.ExternalSource,
	}

	t := s.db.Create(usrProfile)
	if t.Error != nil {
		return nil, t.Error
	}

	return in.UserProfile, nil
}
