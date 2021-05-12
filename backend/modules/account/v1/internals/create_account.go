package account

import (
	"alfred/modules/account/v1/internals/pb"
	"alfred/modules/account/v1/models"
	"context"
	"github.com/hashicorp/go-uuid"
)

func (s *accountInternalServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	account := &models.Account{
		ID:     id,
		Slug:   in.GetAccount().GetSlug(),
		UserID: in.GetAccount().GetUserId(),
	}
	t := s.db.Create(account)
	if t.Error != nil {
		return nil, t.Error
	}
	in.Account.Id = id
	return in.Account, nil
}
