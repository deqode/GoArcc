package external_svc

import (
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *vcsConnectionServer) ListVCSConnection(ctx context.Context, in *pb.ListVCSConnectionRequest) (*pb.ListVCSConnectionResponse, error) {

	var record []models.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId)
	if in.Provider != types.VCSProviders_UNKNOWN {
		chain = chain.Where("provider = ?", in.Provider)
	}
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
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
