package external_svc

import (
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *vcsConnectionServer) GetAccountVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.AccountVCSConnection, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if err := s.ValidateUser(ctx, in.GetAccountId()); err != nil {
		return nil, err
	}
	var record models.VCSConnection
	//todo : by default why there is  and condition
	chain := s.db.Where("account_id = ?", in.AccountId).Where("id = ?", in.Id)
	if err := chain.Find(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	accountVCSConnection := &pb.AccountVCSConnection{
		Id:        record.ID,
		Provider:  record.Provider,
		AccountId: record.AccountID,
		Label:     record.Label,
		UserName:  record.UserName,
	}
	return accountVCSConnection, nil
}
