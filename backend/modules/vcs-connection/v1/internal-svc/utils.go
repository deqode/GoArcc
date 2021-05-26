package internal_svc

import (
	accModel "alfred/modules/account/v1/models"
	"alfred/util/userinfo"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ValidateUserFromContext(ctx context.Context, accountId string, db *gorm.DB) error {
	//Authentication check
	usr := userinfo.FromContext(ctx)
	account := accModel.Account{}
	// select * from account where id  = account id
	if err := db.First(&account, "id = ?", accountId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	//check whether context user id matches with the context user id
	if account.UserID != usr.ID {
		return status.Error(codes.PermissionDenied, "Permission Denied")
	}
	return nil
}
