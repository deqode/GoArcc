package account_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1"
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("Describe:GetAccount", func() {
	var (
		accountServer pb.AccountsServer
	)
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Take this level from config
		Development: false,
	})

	//getting config
	cfgFile, err := config.LoadConfig("config", "/home/deq/Desktop/platform/backend")
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
	//service initialisation
	accountServer = account.NewAccountsServer(fields.db, fields.config, fields.grpcClient)
	Describe("Describe:Categorizing with invalid id", func() {
		Context("Context:When id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.GetAccount(context.Background(), &pb.GetAccountRequest{Id: ""})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})
	})

	Describe("Describe:Categorizing with valid id", func() {
		Context("Context:When id is not empty", func() {
			It("It:Error must be returned", func() {
				acc, _ := accountServer.GetAccount(context.Background(), &pb.GetAccountRequest{Id: "ed521c18-daf0-8516-73b4-681d5cb8fd24"})
				Expect(acc).Should(Equal(acc))
			})
		})
	})
})
