package AccountService

import (
	"alfred/modules/AccountService/v1/models"
	"alfred/modules/AccountService/v1/pb"
	"context"
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AccountService) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	if len(in.Account.UserId) < 3 {
		return nil, status.Error(codes.FailedPrecondition, "UserId is required")
	}
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	account := &models.Account{
		ID:     id,
		Slug:   in.GetAccount().GetSlug(),
		UserId: in.GetAccount().GetUserId(),
	}
	t := s.db.Create(account)
	if t.Error != nil {
		return nil, t.Error
	}
	in.Account.Id = id
	return in.Account, nil
}
