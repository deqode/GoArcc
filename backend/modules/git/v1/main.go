package git

import (
	cadenceAdapter "alfred/background/adapters/cadence"
	background "alfred/background/config"
	"alfred/config"
	"alfred/modules/git/v1/github"
	"alfred/modules/git/v1/pb"
	vcsinternal "alfred/modules/vcs-connection/v1/internals"
	vcsinternalPb "alfred/modules/vcs-connection/v1/internals/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type gitServer struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	vcsInternalServer vcsinternalPb.VCSConnectionInternalServer
	githubService     github.Service
	cadenceConfig     *background.CadenceAppConfig
	cadenceAdapter    *cadenceAdapter.CadenceAdapter
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
	vcsInternalSrv := vcsinternal.NewVCSConnectionInternalServer(db, config, grpcClientConn)
	return &gitServer{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		vcsInternalServer: vcsInternalSrv,
		cadenceAdapter:    cadenceAdapter,
		cadenceConfig:     cadenceConfig,
	}
}
