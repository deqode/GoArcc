package internal_svc

import (
	"alfred/config"
	intPb "alfred/modules/account/v1/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Account Internal Server Configuration
type accountsIntServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

func NewAccountsInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) intPb.AccountInternalServer {
	return &accountsIntServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
