package user_profile_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	pb2 "alfred/modules/account/v1/external-svc/pb"
	. "github.com/onsi/ginkgo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("UpdateUserProfile", func() {
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

	Describe("Update a user profile", func() {
		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
			})
		})

		Context("Get an error when user profile object is nil", func() {
			It("should return nil exception", func() {
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
			})
		})
		Context("Get an error when user profile does not exist", func() {
			It("should return failed precondition error when profile not exist", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("user profile should be updated successfully", func() {
			})
		})

		Context("Update user profile when update mask is correct and id is correct", func() {
			It("check the response of object that user profile is updated or not", func() {
			})
		})
	})
})
