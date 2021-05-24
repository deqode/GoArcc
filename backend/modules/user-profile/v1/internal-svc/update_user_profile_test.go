package internal_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	internal_svc "alfred/modules/user-profile/v1/internal-svc"
	UserProfilepb "alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("UpdateUserProfile", func() {
	var (
		UserProfileServer UserProfilepb.UserProfileInternalServer
		cfg               *config.Config
	)
	BeforeEach(func() {
		//getting config
		cfgFile, err := config.LoadConfig("config", "./../../../../")
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
		UserProfileServer = internal_svc.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Update a user profile", func() {
		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when request object is nil", func() {
			It("should return nil exception", func() {
				_, err := UserProfileServer.UpdateUserProfile(context.Background(), &UserProfilepb.UpdateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile to update is not provided")))
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
