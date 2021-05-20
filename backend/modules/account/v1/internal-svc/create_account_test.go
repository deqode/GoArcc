package internal_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/internal-svc"
	"alfred/modules/account/v1/internal-svc/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("CreateAccount", func() {
	var (
		accountServer pb.AccountInternalServer
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
		accountServer = internal_svc.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Describe:Creating Account Correspond to user", func() {
		Context("Context:When user_id is empty", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(context.Background(), &pb.CreateAccountRequest{Account: &pb.Account{
					Id:     "",
					Slug:   "",
					UserId: "",
				}})
				Expect(err.(pb.CreateAccountRequestValidationError).Cause().(pb.AccountValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Context:When request is nil", func() {
			It("It:Error must be returned", func() {
				_, err := accountServer.CreateAccount(context.Background(), &pb.CreateAccountRequest{Account: nil})
				Expect(err.(pb.CreateAccountRequestValidationError).Reason()).Should(Equal("value is required"))
			})
		})

		Context("Context:Happy Path", func() {
			It("It: should create the account", func() {
				_, err := accountServer.CreateAccount(context.Background(), &pb.CreateAccountRequest{Account: &pb.Account{
					Id:     "",
					Slug:   "",
					UserId: "",
				}})
				Expect(err.(pb.CreateAccountRequestValidationError).Reason()).Should(Equal("value is required"))
			})
		})
	})
})
