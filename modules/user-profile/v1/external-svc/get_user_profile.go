package external_svc

import (
	"context"
	"errors"
	model "github.com/deqode/GoArcc/modules/user-profile/v1/models"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"github.com/deqode/GoArcc/util/userinfo"
	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
)

func (s *userProfilesServer) GetUserProfile(ctx context.Context, empty *empty.Empty) (*pb.UserProfile, error) {
	//user profile model
	var usrProfile model.UserProfile
	usrId := userinfo.FromContext(ctx).ID
	gormDb := s.db
	//ie: Select * from account where id = in.id
	if err := gormDb.First(&usrProfile, "id= ?", usrId).Error; err != nil {
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
