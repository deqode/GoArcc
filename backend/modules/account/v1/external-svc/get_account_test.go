package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	pb2 "alfred/modules/account/v1/pb"
	. "github.com/onsi/ginkgo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("Describe:GetAccount", func() {
	var (
		accountServer pb2.AccountsServer
		cfg           *config.Config
	)
	BeforeEach(func() {
		//getting config
		cfgFile, err := config.LoadConfig("config", "./../../../")
		if err != nil {
			log.Fatal(err)
		}
		cfg, err = config.ParseConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
	})
	JustBeforeEach(func() {
		fields := struct {
			db         *gorm.DB
			config     *config.Config
			grpcClient *grpc.ClientConn
		}{
			db:         db.NewConnection(cfg),
			config:     cfg,
			grpcClient: grpcClient.GetGrpcClientConnection(cfg),
		}
		//service initialisation
		accountServer = external_svc.NewAccountExtServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Get an account", func() {
		By("internal or external call")
		Context("Get an error when id is empty", func() {
			It("it should return validation error", func() {

			})
		})
		Context("Get an error when id is wrong", func() {
			It("should return not found error", func() {

			})
		})

		Context("Get a record when id is provided", func() {
			It("should return requested field in the object", func() {})
		})
	})
})
