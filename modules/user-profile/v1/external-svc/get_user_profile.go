package external_svc

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	model "goarcc/modules/user-profile/v1/models"
	"goarcc/modules/user-profile/v1/pb"
	"goarcc/util/userinfo"
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
