package internal_svc

import (
	model "alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *vcsConnectionIntServer) GetVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.VCSConnection, error) {
	var record model.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId).Where("id = ?", in.Id)
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}
	return nil, nil
}
