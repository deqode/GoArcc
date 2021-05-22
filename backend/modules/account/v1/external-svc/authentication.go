package external_svc

import (
	model "alfred/modules/account/v1/models"
	"alfred/util/userinfo"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *accountExtServer) ValidateUser(ctx context.Context, accountId string) error {
	//Authentication check
	usr := userinfo.FromContext(ctx)
	account := model.Account{}
	// select * from account where id  = account id
	if err := s.db.First(&account, "id = ?", accountId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	//check whether context user id matches with the context user id
	if account.UserID != usr.ID {
		return status.Error(codes.PermissionDenied, "unauthenticated user")
	}
	return nil
}
