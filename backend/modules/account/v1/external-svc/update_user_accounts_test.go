package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	pb2 "alfred/modules/account/v1/external-svc/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("UpdateUserAccounts", func() {
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
		accountServer = external_svc.account.NewAccountsServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Describe:Categorizing with invalid id", func() {
		Context("Context:When account id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.UpdateAccount(context.Background(), &pb2.UpdateAccountRequest{Account: &pb2.Account{
					Id:     "",
					Slug:   "",
					UserId: "",
				}})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Account Id is not provided")))
			})
		})
		Context("Context:When account is nil", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.UpdateAccount(context.Background(), &pb2.UpdateAccountRequest{Account: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Account to update is not provided")))
			})
		})
	})
})
