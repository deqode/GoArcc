package internal_svc

import (
	"context"
	model "github.com/deqode/GoArcc/modules/user-profile/v1/models"
	"github.com/deqode/GoArcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

func (s *userProfileInServer) CreateUserProfile(ctx context.Context, in *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	if in.UserProfile == nil || in == nil {
		return nil, status.Error(codes.FailedPrecondition, "UserProfile not provided")
	}
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
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
	if tx.Error != nil {
		return nil, tx.Error
	}
	return in.UserProfile, nil
}
