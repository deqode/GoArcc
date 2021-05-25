package common

import (
	"context"
	"gorm.io/gorm"
)

func ValidateUser(ctx context.Context, accountId string, db *gorm.DB) error {
	return nil
}
