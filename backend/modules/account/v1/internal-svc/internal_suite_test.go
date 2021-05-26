package internal_svc_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	"alfred/modules/account/v1/internal-svc"
	"alfred/modules/account/v1/pb"
	usrExt "alfred/modules/user-profile/v1/external-svc"
	usrInt "alfred/modules/user-profile/v1/internal-svc"
	usrPb "alfred/modules/user-profile/v1/pb"
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
	AccountServerTest pb.AccountInternalServer
	CtxTest           context.Context
	UsrProfile        *usrPb.UserProfile
)

func TestAccountInternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestAccountInternalSvc Service Suite")
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

	//Int userProfile service initialisation
	userProfileIntServer := usrInt.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)

	//Ext userProfile service initialisation (used only for Creation Tables in DB)
	_ = usrExt.NewUserProfilesServer(fields.db, fields.config, fields.grpcClient)

	// Int Account Service initialisation
	actInServer := internal_svc.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)

	// Ext Account Service initialisation (used only for Creation Tables in DB)
	_ = external_svc.NewAccountExtServer(fields.db, fields.config, fields.grpcClient)

	ctx := context.Background()

	//create UserProfile
	id := "github_" + faker.Username()
	// Create a UserProfile before getting
	res, err := userProfileIntServer.CreateUserProfile(ctx, &usrPb.CreateUserProfileRequest{UserProfile: &usrPb.UserProfile{
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
		log.Fatal(err)
	}
	ui := userinfo.UserInfo{
		ID:          res.Id,
		Email:       res.Email,
		Sub:         res.Sub,
		TokenExpiry: time.Time{},
	}
	ctx = userinfo.NewContext(context.Background(), ui)
	//initialize to global variable here
	CtxTest = ctx
	AccountServerTest = actInServer
	UsrProfile = res
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	AccountServerTest = nil
	CtxTest = nil
	UsrProfile = nil
})
