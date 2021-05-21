package external_svc

import (
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *vcsConnectionServer) ListVCSConnection(ctx context.Context, in *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {

	var record []models.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId)
	if in.Provider != types.VCSProviders_UNKNOWN {
		chain = chain.Where("provider = ?", in.Provider)
	}
	if err := chain.Find(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	var accountVcsConnection []*pb.AccountVCSConnection
	for _, v := range record {
		accountVcsConnection = append(accountVcsConnection, &pb.AccountVCSConnection{
			Id:        v.ID,
			Provider:  v.Provider,
			AccountId: v.AccountID,
			Label:     v.Label,
			UserName:  v.UserName,
		})
	}
	return &pb.ListVCSConnectionResponse{
		VcsConnection: accountVcsConnection,
	}, nil
}
