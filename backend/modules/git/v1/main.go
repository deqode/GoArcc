package git

import (
	"alfred/config"
	"alfred/modules/git/v1/pb"
	vcsinternal "alfred/modules/vcs-connection/v1/internals"
	vcsinternalPb "alfred/modules/vcs-connection/v1/internals/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type gitServer struct {
	db          *gorm.DB
	config      *config.Config
	grpcClient  *grpc.ClientConn
	vcsInServer vcsinternalPb.VCSConnectionInternalServer
}

// NewGitServer todo : AlWays add migration code for best practices
func NewGitServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.GitsServer {

	//initial migration of databases: schema migration
	vcsInSrv := vcsinternal.NewVCSConnectionInternalServer(db, config, grpcClientConn)
	return &gitServer{
		db:          db,
		config:      config,
		grpcClient:  grpcClientConn,
		vcsInServer: vcsInSrv,
	}
}
