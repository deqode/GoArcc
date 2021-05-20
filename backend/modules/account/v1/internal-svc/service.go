package internal_svc

import (
	"alfred/config"
	"alfred/modules/account/v1/controllers"
	intPb "alfred/modules/account/v1/internal-svc/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//Account Internal Server Configuration
type accountsIntServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
	store      controllers.AccountStore
}

func NewAccountsInServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) intPb.AccountInternalServer {
	accountStore := controllers.NewAccountStore(db)
	return &accountsIntServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
		store:      accountStore,
	}
}
