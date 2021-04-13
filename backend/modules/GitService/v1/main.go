package GitService

import (
	"alfred/config"
	"alfred/modules/GitService/v1/pb"
	vcsinternal "alfred/modules/VCSConnectionService/v1/internals"
	vcsinternalPb "alfred/modules/VCSConnectionService/v1/internals/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GitService struct {
	db                *gorm.DB
	config            *config.Config
	grpcClient        *grpc.ClientConn
	vcsInternalServer vcsinternalPb.VCSConnectionInternalServiceServer
}

//todo : AlWays add migration code for best practices
func NewGitService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.GitServiceServer {

	//initial migration of databases: schema migration
	vcsInternalSrv := vcsinternal.NewVCSConnectionInternalService(db, config, grpcClientConn)
	return &GitService{
		db:                db,
		config:            config,
		grpcClient:        grpcClientConn,
		vcsInternalServer: vcsInternalSrv,
	}
}
