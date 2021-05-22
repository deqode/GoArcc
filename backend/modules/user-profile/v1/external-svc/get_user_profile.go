package external_svc

import (
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *userProfilesServer) GetUserProfile(ctx context.Context, in *pb.GetUserProfileRequest) (*pb.UserProfile, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//check user
	if err := s.ValidateUser(ctx, in.GetId()); err != nil {
		return nil, err
	}
	//user profile model
	var usrProfile model.UserProfile
	gormDb := s.db
	//ie: Select * from account where id = in.id
	if err := gormDb.First(&usrProfile, "id= ?", in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &pb.UserProfile{
		Id:             usrProfile.ID,
		Sub:            usrProfile.Sub,
		Name:           usrProfile.Name,
		UserName:       usrProfile.UserName,
		Email:          usrProfile.Email,
		PhoneNumber:    usrProfile.PhoneNumber,
		ExternalSource: usrProfile.Source,
		ProfilePicUrl:  usrProfile.ProfilePicURL,
	}, nil
}
