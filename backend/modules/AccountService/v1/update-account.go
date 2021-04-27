package AccountService

import (
	"alfred/modules/AccountService/v1/models"
	"alfred/modules/AccountService/v1/pb"
	"context"
)

func (s *AccountService) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	account := &models.Account{
		Slug: in.GetAccount().GetSlug(),
	}
	t := s.db.Model(&account).Where("id ="+in.Account.Id, true).Update("slug", in.Account.Slug)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.Account, nil
}
