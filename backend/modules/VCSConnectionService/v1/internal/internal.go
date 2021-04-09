package internal

import (
	"alfred/config"
	"alfred/modules/VCSConnectionService/v1/internal/pb"
	"alfred/modules/VCSConnectionService/v1/models"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type VCSConnectionService struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

func NewVCSConnectionInternalService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionInternalServiceServer {
	return &VCSConnectionService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}

func (s *VCSConnectionService) GetVCSConnection(ctx context.Context, in *pb.GetVCSConnectionRequest) (*pb.VCSConnection, error) {
	record := models.VCSConnection{}
	result := s.db.Where("provider = ?", in.Provider).Find(&record)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "No Record Found")
	}

	VCS := &pb.VCSConnection{
		Id:                 record.ID,
		Provider:           1,
		ConnectionId:       record.ConnectionId,
		AccessToken:        record.AccessToken,
		RefreshToken:       record.RefreshToken,
		AccessTokenExpiry:  nil,
		RefreshTokenExpiry: nil,
		Revoked:            false,
		AccountId:          "",
	}
	return VCS, nil
}
func (s *VCSConnectionService) CreateVCSConnection(ctx context.Context, in *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	aTEP, err := ptypes.Timestamp(in.VcsConnection.AccessTokenExpiry)
	if err != nil {
		return nil, err
	}
	rTEP, err := ptypes.Timestamp(in.VcsConnection.RefreshTokenExpiry)
	if err != nil {
		return nil, err
	}
	VCSModel := &models.VCSConnection{
		ID:                 id,
		Provider:           in.VcsConnection.Provider,
		ConnectionId:       in.VcsConnection.ConnectionId,
		AccessToken:        in.VcsConnection.AccessToken,
		RefreshToken:       in.VcsConnection.RefreshToken,
		AccessTokenExpiry:  &aTEP,
		RefreshTokenExpiry: &rTEP,
		Revoked:            false,
		AccountId:          in.VcsConnection.AccountId,
	}

	t := s.db.Create(VCSModel)
	if t.Error != nil {
		return nil, t.Error
	}
	return in.VcsConnection, nil
}
func (s *VCSConnectionService) RevokeVCSToken(context.Context, *pb.RevokeVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
func (s *VCSConnectionService) RenewVCSToken(context.Context, *pb.RenewVCSTokenRequest) (*empty.Empty, error) {
	return nil, nil
}
