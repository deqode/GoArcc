package external_svc

import (
	"context"
	"errors"
	model "goarcc/modules/user-profile/v1/models"
	"goarcc/modules/user-profile/v1/pb"
	"gorm.io/gorm"
)

func (s *userProfilesServer) GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//user profile model
	usrProfile := model.UserProfile{}
	gormDb := s.db
	//ie: Select * from account where id = in.id
	if err := gormDb.First(&usrProfile, "sub= ?", in.Sub).Error; err != nil {
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
