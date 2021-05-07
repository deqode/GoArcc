package internals

import (
	"alfred/config"
	"alfred/modules/VCSConnection/v1/internals/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionInternalServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

func NewVCSConnectionInternalServer(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.VCSConnectionInternalServer {
	return &vcsConnectionInternalServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
