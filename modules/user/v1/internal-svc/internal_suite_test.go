package internal_svc_test

import (
	"context"
	"github.com/deqode/GoArcc/client/grpcClient"
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/db"
	"github.com/deqode/GoArcc/logger"
	"github.com/deqode/GoArcc/modules/user/v1/internal-svc"
	"github.com/deqode/GoArcc/modules/user/v1/pb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"
)

var (
	UserServerIntTest pb.UserInternalServer
	CtxTest           context.Context
	Usr               *pb.User
)

func TestUserInternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestUser Internal Service Suite")
}

// This block will run only once
var _ = BeforeSuite(func() {
	//getting config
	cfgFile, err := config.LoadConfig("config", config.GetConfigDirectory())
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
	//Int service initialisation
	userIntServer := internal_svc.NewUserInServer(fields.db, fields.config)

	//initialize to global variable here
	UserServerIntTest = userIntServer
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	Usr = nil
	UserServerIntTest = nil
})
