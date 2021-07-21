package external_svc_test

import (
	"github.com/deqode/GoArcc/client/grpcClient"
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/db"
	"github.com/deqode/GoArcc/logger"
	"github.com/deqode/GoArcc/modules/authentication/v1/external-svc"
	"github.com/deqode/GoArcc/modules/authentication/v1/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	AuthServerTest pb.AuthenticationsServer
)

func TestAuthenticationExt(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestAuthenticationExt Service Suite")
}

// Before Suite Run only once
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

	authServer := external_svc.NewAuthenticationServer(fields.db, fields.config, fields.grpcClient)
	_, err = external_svc.NewAuthenticator(fields.config)
	if err != nil {
		log.Fatal(err)
	}

	AuthServerTest = authServer
})

// nil all global variables
var _ = AfterSuite(func() {
	AuthServerTest = nil
})
