package internal_svc

import (
	model "alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *vcsConnectionIntServer) GetVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.VCSConnection, error) {
	// Request Validation
	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request is nil")
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// User Validation
	if err := ValidateUserFromContext(ctx, in.AccountId, s.db); err != nil {
		return nil, err
	}
	var record model.VCSConnection
	chain := s.db.Where("account_id = ?", in.AccountId)
	if in.Provider != types.VCSProviders_UNKNOWN {
		chain = chain.Where("provider= ?", in.GetProvider())
	}
	tx := chain.Find(&record)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0{
		return nil,gorm.ErrRecordNotFound
	}
	aTokenExpiry, err := ptypes.TimestampProto(*record.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}
	rTokenExpiry, err := ptypes.TimestampProto(*record.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}
	return &pb.VCSConnection{
		Id:                 record.ID,
		Label:              record.Label,
		Provider:           record.Provider,
		ConnectionId:       record.ConnectionID,
		AccessToken:        record.AccessToken,
		RefreshToken:       record.RefreshToken,
		AccessTokenExpiry:  aTokenExpiry,
		RefreshTokenExpiry: rTokenExpiry,
		Revoked:            record.Revoked,
		AccountId:          record.AccountID,
		UserName:           record.UserName,
	}, nil
}
