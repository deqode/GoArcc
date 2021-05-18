package account_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1"
	"alfred/modules/account/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("Describe:Delete account test", func() {
	var (
		accountServer pb.AccountsServer
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
		accountServer = account.NewAccountsServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Describe:Categorizing with invalid id", func() {
		Context("Context:When id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.DeleteAccount(context.Background(), &pb.DeleteAccountRequest{Id: ""})
				Expect(err.(pb.DeleteAccountRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
	})
})
