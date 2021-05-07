package internals

import (
	"alfred/modules/VCSConnection/v1/internals/pb"
	"alfred/modules/VCSConnection/v1/models"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *vcsConnectionInternalServer) GetVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.VCSConnection, error) {
	record := models.VCSConnection{}
	chain := s.db.Where("provider = ?", in.Provider)
	if in.AccountId != "" {
		chain = chain.Where("account_id = ?", in.AccountId)
	}
	result := chain.Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}

	VCS := &pb.VCSConnection{
		Id:                 record.ID,
		Provider:           in.Provider,
		ConnectionId:       record.ConnectionID,
		AccessToken:        record.AccessToken,
		RefreshToken:       record.RefreshToken,
		AccessTokenExpiry:  nil,
		RefreshTokenExpiry: nil,
		Revoked:            false,
		AccountId:          in.AccountId,
		UserName:           record.UserName,
		Label:              record.Label,
	}
	return VCS, nil
}
