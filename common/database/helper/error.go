package helper

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ValidateResult(tx *gorm.DB) error {
	if tx == nil {
		return status.Error(codes.Internal, "something went wrong")
	}
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return tx.Error
		}
		return status.Error(codes.Internal, tx.Error.Error())
	}
	return nil
}
