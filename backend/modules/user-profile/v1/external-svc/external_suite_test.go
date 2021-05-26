package external_svc_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	external_svc "alfred/modules/user-profile/v1/external-svc"
	"alfred/modules/user-profile/v1/internal-svc"
	"alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"alfred/util/userinfo"
	"context"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

var (
	UserProfileServerTest pb.UserProfilesServer
	CtxTest               context.Context
	UsrProfile            *pb.UserProfile
)

func TestUserProfileExternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestUserProfile External Service Suite")
}

// This block will run only once
var _ = BeforeSuite(func() {
	//getting config
	cfgFile, err := config.LoadConfig("config", "./../../../../")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	fields := struct {
		db         *gorm.DB
		config     *config.Config
		grpcClient *grpc.ClientConn
	}{
		db:         db.NewConnection(cfg),
		config:     cfg,
		grpcClient: grpcClient.GetGrpcClientConnection(cfg),
	}
	//Ext service initialisation
	userProfileServer := external_svc.NewUserProfilesServer(fields.db, fields.config, fields.grpcClient)
	//Int service initialisation
	userProfileIntServer := internal_svc.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)

	id := "github_" + faker.Username()
	// Create a UserProfile before getting
	res, err := userProfileIntServer.CreateUserProfile(CtxTest, &pb.CreateUserProfileRequest{UserProfile: &pb.UserProfile{
		Id:             id,
		Sub:            id,
		Name:           faker.Name(),
		UserName:       faker.Username(),
		Email:          faker.Email(),
		PhoneNumber:    faker.Phonenumber(),
		ExternalSource: types.VCSProviders_GITHUB,
		ProfilePicUrl:  faker.URL(),
		TokenValidTill: nil,
	}})
	if err != nil {
		return
	}
	ui := userinfo.UserInfo{
		ID:          res.Id,
		Email:       res.Email,
		Sub:         res.Sub,
		TokenExpiry: time.Time{},
	}

	//initialize to global variable here
	CtxTest = userinfo.NewContext(context.Background(), ui)
	UsrProfile = res
	UserProfileServerTest = userProfileServer
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	UsrProfile = nil
	UserProfileServerTest = nil
})
