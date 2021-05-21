package external_svc

import (
	"alfred/config"
	internal_svc "alfred/modules/vcs-connection/v1/internal-svc"
	"alfred/modules/vcs-connection/v1/models"
	"alfred/modules/vcs-connection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type vcsConnectionServer struct {
	db                     *gorm.DB
	config                 *config.Config
	grpcClient             *grpc.ClientConn
	vcsConnectionIntServer pb.VCSConnectionInternalServer
}

// NewVCSConnectionServer todo : AlWays add migration code for best practices
func NewVCSConnectionServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.VCSConnectionsServer {
	//initial migration of databases: schema migration
	models.InitialMigrationVCSConnection(db)
	vcsConnectionIntSrv := internal_svc.NewVCSConnectionIntServer(db, config, grpcClientConn)
	return &vcsConnectionServer{
		db:                     db,
		config:                 config,
		grpcClient:             grpcClientConn,
		vcsConnectionIntServer: vcsConnectionIntSrv,
	}
}
