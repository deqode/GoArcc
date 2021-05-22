package external_svc

import (
	accountModel "alfred/modules/account/v1/models"
	"alfred/util/userinfo"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// ValidateUser : ValidateUser
func (s *vcsConnectionServer) ValidateUser(ctx context.Context, accountId string) error {
	//Authentication check
	usr := userinfo.FromContext(ctx)
	account := accountModel.Account{}
	// select * from account where id  = account id
	if err := s.db.First(&account, "id = ?", accountId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	if account.UserID != usr.ID {
		return status.Error(codes.PermissionDenied, "unauthenticated user")
	}
	return nil
}
