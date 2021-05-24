package external_svc

import (
	userModel "alfred/modules/user-profile/v1/models"
	"alfred/util/userinfo"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *userProfilesServer) ValidateUser(ctx context.Context, userId string) error {
	//Authentication check
	usr := userinfo.FromContext(ctx)
	userModel := userModel.UserProfile{}
	// select * from account where id  = account id
	if err := s.db.First(&userModel, "id = ?", userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	//check whether context user id matches with the  user id
	if userModel.ID != usr.ID {
		return status.Error(codes.PermissionDenied, "unauthenticated user")
	}
	return nil
}
