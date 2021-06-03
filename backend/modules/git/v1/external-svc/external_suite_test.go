package external_svc_test

import (
	"alfred.sh/common/logger"
	background "alfred/background/cmd"
	"alfred/client/grpcClient"
	"alfred/config"
	"alfred/db"
	accInt "alfred/modules/account/v1/internal-svc"
	accPb "alfred/modules/account/v1/pb"
	"alfred/modules/git/v1/external-svc"
	"alfred/modules/git/v1/pb"
	usrInt "alfred/modules/user-profile/v1/internal-svc"
	usrPb "alfred/modules/user-profile/v1/pb"
	vcsInt "alfred/modules/vcs-connection/v1/internal-svc"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
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

func TestGitExternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestGitExternalSvc Suite")
}

var (
	CtxTest          context.Context
	VcsIntServerTest vcsPb.VCSConnectionInternalServer
	GitExtServerTest pb.GitsServer
	VCSTest          *vcsPb.VCSConnection
)

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

	//Ext Git service initialisation
	cadConf, cadAdapt := background.CadenceSetup(fields.config)
	gitExtServer := external_svc.NewGitServer(fields.db, fields.config, fields.grpcClient, cadConf, cadAdapt)

	//Int userProfile service initialisation
	userProfileIntServer := usrInt.NewUserProfileInServer(fields.db, fields.config, fields.grpcClient)

	// Int Account Service initialisation
	actInServer := accInt.NewAccountsInServer(fields.db, fields.config, fields.grpcClient)

	// Int VCS Server initialisation
	vcsIntServer := vcsInt.NewVCSConnectionIntServer(fields.db, fields.config, fields.grpcClient)

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
	vcs := &vcsPb.VCSConnection{
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
	vcs, err = vcsIntServer.CreateVCSConnection(ctx, &vcsPb.CreateVCSConnectionRequest{
		VcsConnection: vcs,
	})
	if err != nil {
		log.Fatal(err)
	}
	//initialize to global variable here
	CtxTest = ctx
	VcsIntServerTest = vcsIntServer
	GitExtServerTest = gitExtServer
	VCSTest = vcs
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	VcsIntServerTest = nil
	VCSTest = nil
})
