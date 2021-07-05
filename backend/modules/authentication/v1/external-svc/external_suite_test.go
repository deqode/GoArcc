package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/authentication/v1/external-svc"
	"alfred/modules/authentication/v1/pb"
	"alfred/util/logger"
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
