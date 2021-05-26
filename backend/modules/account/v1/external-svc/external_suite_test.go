package external_svc_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	"alfred/modules/account/v1/external-svc"
	"alfred/modules/account/v1/internal-svc"
	"alfred/modules/account/v1/pb"
	usrInt "alfred/modules/user-profile/v1/internal-svc"
	usrPb "alfred/modules/user-profile/v1/pb"
	"alfred/protos/types"
	"alfred/util/userinfo"
	"context"
	"github.com/bxcodec/faker/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	AccountServerTest      pb.AccountsServer
	AccountServerInTest    pb.AccountInternalServer
	CtxTest                context.Context
	Account                *pb.Account
	UsrProfileInServerTest usrPb.UserProfileInternalServer
)

func TestAccountExternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestAccountExternal Service Suite")
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

	// Int Account Service initialisation
	actInServer := internal_svc.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)

	// Ext Account Service initialisation
	actExtServer := external_svc.NewAccountExtServer(fields.db, fields.config, fields.grpcClient)

	id := "github_" + faker.Username()
	ctx := context.Background()
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
		return
	}
	ui := userinfo.UserInfo{
		ID:          res.Id,
		Email:       res.Email,
		Sub:         res.Sub,
		TokenExpiry: time.Time{},
	}
	ctx = userinfo.NewContext(ctx, ui)
	// Create A Account
	acc, err := actInServer.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &pb.Account{
		Slug:   res.Sub + "_" + res.UserName,
		UserId: res.Id,
	}})
	if err != nil {
		return
	}

	//initialize to global variable here
	CtxTest = ctx
	AccountServerTest = actExtServer
	AccountServerInTest = actInServer
	UsrProfileInServerTest = userProfileIntServer
	Account = acc
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	AccountServerTest = nil
	AccountServerInTest = nil
	UsrProfileInServerTest = nil
	Account = nil
})
