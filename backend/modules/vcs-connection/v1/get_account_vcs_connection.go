package vcs_connection

import (
	models2 "alfred/modules/vcs-connection/v1/internal/models"
	"alfred/modules/vcs-connection/v1/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *vcsConnectionServer) GetAccountVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.AccountVCSConnection, error) {
	var record models2.VCSConnection
	//todo : by default why there is  and condition
	chain := s.db.Where("account_id = ?", in.AccountId).Where("id = ?", in.Id)
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
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
