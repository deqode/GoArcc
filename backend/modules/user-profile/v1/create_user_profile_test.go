package user_profile_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	userProfile "alfred/modules/user-profile/v1"
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
		UserProfileServer = userProfile.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Describe:Categorizing with invalid user_id", func() {
		Context("Context:When sub is empty", func() {
			It("It:Error must be returned", func() {
				request := Profile
				request.Sub = ""
				_, err := UserProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: request})
				Expect(err.(pb.CreateUserProfileRequestValidationError).Cause().(pb.UserProfileValidationError).Field()).Should(Equal("Sub"))
			})
		})
		Context("Context:When request is nil", func() {
			It("It:Error must be returned", func() {
				_, err := UserProfileServer.CreateUserProfile(context.Background(), &pb.CreateUserProfileRequest{UserProfile: nil})
				Expect(err).Should(Equal(status.Error(codes.FailedPrecondition, "UserProfile not provided")))
			})
		})
	})
})
