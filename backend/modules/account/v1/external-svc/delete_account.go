package external_svc

import (
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"alfred/util/userinfo"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// DeleteAccount : Will Delete account  with the given id. if record not found it will give error
func (s accountExtServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	//request validation
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	if err := userinfo.ValidateUser(ctx , &model.Account{} , s.db) ; err != nil {
		return nil , err
	}
	//by default it will delete with primary key.
	// ie: Delete From Account where id = in.id
	if err := s.db.Where("id = ?", in.Id).Delete(&model.Account{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &empty.Empty{}, nil
}



