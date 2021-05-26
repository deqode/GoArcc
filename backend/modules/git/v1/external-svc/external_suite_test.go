package external_svc_test

import (
	"alfred.sh/common/logger"
	"go.uber.org/zap"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGitExternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestGitExternalSvc Suite")
}

var _ = BeforeSuite(func() {
	//getting config
	//cfgFile, err := config.LoadConfig("config", "./../../../../")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//cfg, err := config.ParseConfig(cfgFile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fields := struct {
	//	db         *gorm.DB
	//	config     *config.Config
	//	grpcClient *grpc.ClientConn
	//}{
	//	db:         db.NewConnection(cfg),
	//	config:     cfg,
	//	grpcClient: grpcClient.GetGrpcClientConnection(cfg),
	//}
	//
	////Ext Git service initialisation
	//cadConf, cadAdapt := background.CadenceSetup(fields.config)
	//_ = external_svc.NewGitServer(fields.db, fields.config, fields.grpcClient,cadConf, cadAdapt)

})

var _ = AfterSuite(func() {

})
