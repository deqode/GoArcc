package git

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/git/v1/pb"
	vcs "alfred/modules/vcs-connection/v1"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type gitServer struct {
	db             *gorm.DB
	config         *config.Config
	grpcClient     *grpc.ClientConn
	vcsInServer    vcsPb.VCSConnectionsServer
	cadenceConfig  *background.CadenceAppConfig
	cadenceAdapter *cadenceAdapter.CadenceAdapter
}

// NewGitServer todo : AlWays add migration code for best practices
func NewGitServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
	cadenceConfig *background.CadenceAppConfig,
	cadenceAdapter *cadenceAdapter.CadenceAdapter,
) pb.GitsServer {

	//initial migration of databases: schema migration
	vcsInSrv := vcs.NewVCSConnectionServer(db, config, grpcClientConn)
	return &gitServer{
		db:             db,
		config:         config,
		grpcClient:     grpcClientConn,
		vcsInServer:    vcsInSrv,
		cadenceAdapter: cadenceAdapter,
		cadenceConfig:  cadenceConfig,
	}
}
