package internal_svc_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/internal-svc"
	"alfred/modules/account/v1/pb"
	usrInt "alfred/modules/user-profile/v1/internal-svc"
	usrPb "alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"
)

var (
	AccountServerTest        pb.AccountInternalServer
	CtxTest                  context.Context
	UserProfileIntServerTest usrPb.UserProfileInternalServer
)

func TestAccountInternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestAccountInternalSvc Service Suite")
}

// This block will run only once
var _ = BeforeSuite(func() {
	//getting config
	cfgFile, err := config.LoadConfig("config", "./../../../../")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	fields := struct {
		db         *gorm.DB
		config     *config.Config
		grpcClient *grpc.ClientConn
	}{
		db:         db.NewConnection(cfg),
		config:     cfg,
		grpcClient: grpcClient.GetGrpcClientConnection(cfg),
	}

	//Int userProfile service initialisation
	userProfileIntServer := usrInt.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)

	// Int Account Service initialisation
	actInServer := internal_svc.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)

	ctx := context.Background()

	//initialize to global variable here
	CtxTest = ctx
	AccountServerTest = actInServer
	UserProfileIntServerTest = userProfileIntServer
})
