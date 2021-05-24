package external_svc_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	external_svc "alfred/modules/user-profile/v1/external-svc"
	pb "alfred/modules/user-profile/v1/pb"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

var _ = Describe("GetUserProfile", func() {
	var (
		UserProfileServer pb.UserProfilesServer
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
		UserProfileServer = external_svc.NewUserProfilesServer(fields.db, fields.config, fields.grpcClient)
	})

	Describe("Get a user-profile", func() {
		By("internal or external call")
		Context("Get an error when id is empty", func() {
			It("it should return validation error", func() {
				_, err := UserProfileServer.GetUserProfileBySub(context.Background(), &pb.GetUserProfileBySubRequest{Sub: ""})
				Expect(err.(pb.GetUserProfileBySubRequestValidationError).Reason()).Should(Equal("value length must be at least 3 runes"))
			})
		})
		Context("Get an error when sub_id is wrong", func() {
			It("should return not found error", func() {
				_, err := UserProfileServer.GetUserProfileBySub(context.Background(), &pb.GetUserProfileBySubRequest{Sub: "fji5895"})
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})
		})

		Context("Get a record when sub_id is provided", func() {
			It("should return requested field in the object", func() {
				//_, err := UserProfileServer.GetUserProfileBySub(context.Background(), &pb.GetUserProfileBySubRequest{Sub: "github_1234"})
				//Expect(err).Should(BeNil())
			})
		})
	})
})
