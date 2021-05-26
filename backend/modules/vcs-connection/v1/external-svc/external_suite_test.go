package external_svc_test

import (
	"alfred.sh/common/logger"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	accInt "alfred/modules/account/v1/internal-svc"
	accPb "alfred/modules/account/v1/pb"
	usrInt "alfred/modules/user-profile/v1/internal-svc"
	usrPb "alfred/modules/user-profile/v1/pb"
	"alfred/modules/vcs-connection/v1/external-svc"
	"alfred/modules/vcs-connection/v1/internal-svc"
	"alfred/modules/vcs-connection/v1/pb"
	"alfred/protos/types"
	"alfred/util/userinfo"
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVCSConnectionExt(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestVCSConnectionExt Service Suite")
}

var (
	CtxTest       context.Context
	VcsServerTest pb.VCSConnectionsServer
	VCSTest       *pb.VCSConnection
)

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
	actInServer := accInt.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)

	// Ext VCS Server initialisation
	vcsExtServer := external_svc.NewVCSConnectionServer(fields.db, fields.config, fields.grpcClient)
	// Int VCS Server initialisation
	vcsIntServer := internal_svc.NewVCSConnectionIntServer(fields.db, fields.config, fields.grpcClient)
	ctx := context.Background()
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
	ctx = userinfo.NewContext(ctx, ui)
	// Create A Account
	acc, err := actInServer.CreateAccount(ctx, &accPb.CreateAccountRequest{Account: &accPb.Account{
		Slug:   res.Sub + "_" + res.UserName,
		UserId: res.Id,
	}})
	if err != nil {
		log.Fatal(err)
	}

	time, err := ptypes.TimestampProto(time.Now().UTC().AddDate(1, 11, 0))
	if err != nil {
		log.Fatal(err)
	}
	//generate VCS from faker data
	vcs := &pb.VCSConnection{
		Id:                 faker.UUIDHyphenated(),
		Label:              faker.FirstName(),
		Provider:           types.VCSProviders_GITHUB,
		ConnectionId:       faker.UUIDDigit(),
		AccessToken:        faker.Password(),
		RefreshToken:       "",
		AccessTokenExpiry:  time,
		RefreshTokenExpiry: time,
		Revoked:            false,
		AccountId:          acc.Id,
		UserName:           faker.Username(),
	}

	// create A VCS Connection as well
	vcs, err = vcsIntServer.CreateVCSConnection(ctx, &pb.CreateVCSConnectionRequest{
		VcsConnection: vcs,
	})
	if err != nil {
		log.Fatal(err)
	}
	//initialize to global variable here
	CtxTest = ctx
	VcsServerTest = vcsExtServer
	VCSTest = vcs
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	VcsServerTest = nil
	VCSTest = nil
})
