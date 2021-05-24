package internal_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/user-profile/v1/internal-svc"
	"alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("DeleteUserProfile", func() {
	var (
		UserProfileServer pb.UserProfileInternalServer
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

	Describe("Deleting a user profile", func() {
		//Negative Cases
		By("By an internal RPC Call")
		Context("Get an error when no user_id provided", func() {
			It("Should return validation error", func() {
				_, err := UserProfileServer.DeleteUserProfile(context.Background(), &pb.DeleteUserProfileRequest{Id: ""})
				Expect(err.(pb.DeleteUserProfileRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})

		Context("Get an error when request is nil", func() {
			It("should return error ", func() {
				_, err := UserProfileServer.DeleteUserProfile(context.Background(), nil)
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "Request can't be nil")))
			})
		})

		Context("Return nil when wrong user_id provided", func() {
			It("should not perform any action in DB and return nil", func() {
				_, err := UserProfileServer.DeleteUserProfile(context.Background(), &pb.DeleteUserProfileRequest{Id: "fejfoeiw0r8290r420"})
				Expect(err).To(Not(BeNil()), "Error")
			})
		})

		//Positive Test Cases
		Context("Return confirmation when record soft deleted in DB", func() {
			It("should return a boolean value", func() {
				//_, err := UserProfileServer.DeleteUserProfile(context.Background(), &pb.DeleteUserProfileRequest{Id: "github_09292"})
				//Expect(err).To(BeNil(), "Error")
			})
		})
	})

})
