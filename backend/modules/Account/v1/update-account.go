package Account

import (
	"alfred/modules/Account/v1/models"
	"alfred/modules/Account/v1/pb"
	"context"
)

func (s *accountsServer) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	account := &models.Account{
		Slug: in.GetAccount().GetSlug(),
	}
	t := s.db.Model(&account).Where("id ="+in.Account.Id, true).Update("slug", in.Account.Slug)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.Account, nil
}
