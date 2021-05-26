package external_svc

import (
	model "alfred/modules/account/v1/models"
	pb "alfred/modules/account/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//todo -
func (s accountExtServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request is nil")
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}
	account := model.Account{
		ID:   in.Id,
		Slug: in.Slug,
	}
	// Omit Exclude fields from updating
	tx := s.db.Model(&account).Omit("id", "user_id").Updates(&account)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(codes.Internal, "Record not updated")
	}
	return &pb.Account{
		Id:     account.ID,
		Slug:   account.Slug,
		UserId: account.UserID,
	}, nil
}
