package external_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
)

func (s *userProfilesServer) GetUserProfileBySub(ctx context.Context, in *pb.GetUserProfileBySubRequest) (*pb.UserProfile, error) {
	// TODO - add ctx validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//user profile model
	usrProfile := model.UserProfile{}
	gormDb := s.db
	//ie: Select * from account where id = in.id
	tx := gormDb.Where(&usrProfile, in.Sub)
	if err := databaseHelper.ValidateResult(tx); err != nil {
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
