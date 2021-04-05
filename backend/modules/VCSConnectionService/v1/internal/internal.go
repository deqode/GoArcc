package internal

import (
	"alfred/config"
	"alfred/modules/VCSConnectionService/v1/internal/pb"
	"alfred/modules/VCSConnectionService/v1/models"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type VCSConnectionService struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

//todo : AlWays add migration code for best practices
func NewVCSConnectionService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionServiceServer {

	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	return &VCSConnectionService{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}

func (s *VCSConnectionService) GetVCSConnection(context.Context, *pb.GetVCSConnectionRequest) (*pb.VCSConnection, error){
	return nil, nil
}
func (s *VCSConnectionService) CreateVCSConnection(context.Context, *pb.CreateVCSConnectionRequest) (*pb.VCSConnection, error){
	return nil, nil
}
func (s *VCSConnectionService) RevokeVCSToken(context.Context, *pb.RevokeVCSTokenRequest) (*empty.Empty, error){
	return nil, nil
}
func (s *VCSConnectionService) RenewVCSToken(context.Context, *pb.RenewVCSTokenRequest) (*empty.Empty, error){
	return nil, nil
}
