package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	"alfred/modules/account/v1/pb"
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
		accountServer = external_svc.NewAccountExtServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Update an account", func() {
		//Negative Test Cases
		By("By a internal or external RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
				_, err := accountServer.UpdateAccount(context.Background(), &pb.UpdateAccountRequest{Account: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Account to update is not provided")))
			})
		})

		Context("Get an error when update mask is incorrect ", func() {
			It("should return failed precondition error", func() {
			})
		})

		Context("Get an error when update mask contain id ", func() {
			It("should return failed precondition error", func() {
			})
		})
		Context("Get an error when id is incorrect", func() {
			It("should return failed precondition error", func() {
				_, err := accountServer.UpdateAccount(context.Background(), &pb.UpdateAccountRequest{Account: &pb.Account{
					Id:     "",
					Slug:   "",
					UserId: "",
				}})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Account Id is not provided")))
			})
		})
		Context("Get an error when account does not exist", func() {
			It("should return failed precondition error when profile not exist", func() {
			})
		})

		Context("Update account when update mask is correct and id is correct", func() {
			It("user Account should be updated successfully", func() {
			})
		})

		Context("Update user account when update mask is correct and id is correct", func() {
			It("check the response of object that account is updated or not", func() {
			})
		})

	})
})
