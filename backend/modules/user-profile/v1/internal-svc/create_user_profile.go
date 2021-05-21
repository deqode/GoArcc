package internal_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/user-profile/v1/models"
	"alfred/modules/user-profile/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

func (s *userProfileInServer) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if in.UserProfile == nil {
		return nil, status.Error(codes.FailedPrecondition, "UserProfile not provided")
	}
	//prepare insert object
	UserProfileModel := &model.UserProfile{
		ID:            in.GetUserProfile().GetId(),
		Name:          in.GetUserProfile().GetName(),
		UserName:      in.GetUserProfile().GetUserName(),
		Email:         in.GetUserProfile().GetEmail(),
		PhoneNumber:   in.GetUserProfile().GetPhoneNumber(),
		Sub:           in.GetUserProfile().GetId(),
		ProfilePicURL: in.GetUserProfile().GetProfilePicUrl(),
		Source:        in.GetUserProfile().GetExternalSource(),
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
	//insert into db
	gormDb := s.db
	tx := gormDb.Create(UserProfileModel)
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	return &pb.UserProfile{
		Id: UserProfileModel.ID,
	}, nil
}
