package external_svc

import (
	model "alfred/modules/account/v1/models"
	"alfred/modules/account/v1/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteAccount : Will Delete account  with the given id. if record not found it will give error
func (s accountExtServer) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*empty.Empty, error) {
	//request validation
	if err := in.Validate(); err != nil {
		return nil, err
	}

	tx := s.db.Delete(&model.Account{ID: in.Id})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.Internal, "Record not deleted")
	}
	return &empty.Empty{}, nil
}
