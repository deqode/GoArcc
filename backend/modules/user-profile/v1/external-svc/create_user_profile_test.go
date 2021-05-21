package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/user-profile/v1/external-svc"
	"alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("CreateUserProfile", func() {
	var (
		Profile           *pb.UserProfile
		UserProfileServer pb.UserProfileInternalServer
		cfg               *config.Config
	)
	BeforeEach(func() {
		Profile = &pb.UserProfile{
			Id:             "",
			Sub:            "github_09292",
			Name:           "Shivang Saxena",
			UserName:       "shivang-saxena",
			Email:          "ssaxena@deqode.com",
			PhoneNumber:    "9838392392",
			ExternalSource: types.VCSProviders_GITHUB,
			ProfilePicUrl:  "https://images.google.com",
			TokenValidTill: nil,
		}
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
		UserProfileServer = external_svc.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Creating an user profile", func() {

		//Negative Test Cases
		By("By a internal RPC Call")
		Context("Get an error when nil User provided", func() {
			It("should return nil exception", func() {
				_, err := UserProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile not provided")))
			})
		})

		Context("Create a user when subject is empty", func() {
			It("It should return validation error", func() {
				request := Profile
				request.Sub = ""
				_, err := UserProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				Expect(err.(pb.CreateUserProfileRequestValidationError).Cause().(pb.UserProfileValidationError).Field()).Should(Equal("Sub"))
			})
		})

		Context("Get an error when user created with wrong email", func() {
			It("if email is not provided then ignore because its a non required field", func() {})
			It("Should return validation error when email is provided but not formatted", func() {})
		})

		Context("Get an error when user created with wrong Phone", func() {
			It("if phone is not provided then ignore because its a non required field", func() {})
			It("Should return validation error when phone number is not in proper", func() {})
		})

		Context("Get an error when name of user not provided", func() {
			It("if name is nil", func() {})
			It("if name length exceed maximum upto 100 character", func() {})
		})

		Context("Get an error when user-name of user not provided", func() {
			It("if user-name is empty return an error", func() {})
			It("if user-name length exceed maximum upto 100 character return error", func() {})
		})

		Context("Return proper error when user is creating from unknown source", func() {
			It("should return error if user is from unknown source", func() {})
		})

		Context("Get an error when wrong profile-address provided", func() {
			It("if not a valid url return error", func() {})

		})
		Context("Get an error when user create with already existing sub", func() {
			It("should return already exists error", func() {})
		})

		//Positive Test Cases
		Context("Create a user profile", func() {
			It("should return user_id as uuid", func() {
			})
		})

	})
})
