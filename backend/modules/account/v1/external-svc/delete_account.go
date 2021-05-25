package external_svc

import (
	databaseHelper "alfred.sh/common/database/helper"
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

// DeleteAccount : Will Delete account  with the given id. if record not found it will give error
func (s accountExtServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	//request validation
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	tx := s.db.Where("id = ?", in.Id).Delete(&model.Account{})
	if err := databaseHelper.ValidateResult(tx); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
