package external_svc_test

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/deqode/GoArcc/config"
	"github.com/deqode/GoArcc/db"
	"github.com/deqode/GoArcc/logger"
	external_svc "github.com/deqode/GoArcc/modules/user/v1/external-svc"
	"github.com/deqode/GoArcc/modules/user/v1/internal-svc"
	"github.com/deqode/GoArcc/modules/user/v1/pb"
	"github.com/deqode/GoArcc/util/userinfo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

var (
	UserServerTest pb.UsersServer
	CtxTest        context.Context
	Usr            *pb.User
)

func TestUserProfileExternalSvc(t *testing.T) {
	//now init logger
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestUser External Service Suite")
}

// This block will run only once
var _ = BeforeSuite(func() {
	//getting config
	cfgFile, err := config.LoadConfig("config", config.GetConfigDirectory())
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	fields := struct {
		db     *gorm.DB
		config *config.Config
	}{
		db:     db.NewConnection(cfg),
		config: cfg,
	}
	//Ext service initialisation
	userServer := external_svc.NewUserServer(fields.db, fields.config)
	//Int service initialisation
	userIntServer := internal_svc.NewUserInServer(fields.db, fields.config)

	id := "github_" + faker.Username()
	// Create a UserProfile before getting
	res, err := userIntServer.CreateUser(CtxTest, &pb.CreateUserRequest{User: &pb.User{
		Id:            id,
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		UserName:      faker.Username(),
		Email:         faker.Email(),
		PhoneNumber:   faker.Phonenumber(),
		ProfilePicUrl: faker.URL(),
	}})
	if err != nil {
		return
	}
	ui := userinfo.UserInfo{
		ID:          res.Id,
		Email:       res.Email,
		TokenExpiry: time.Time{},
	}

	//initialize to global variable here
	CtxTest = userinfo.NewContext(context.Background(), ui)
	Usr = res
	UserServerTest = userServer
})

// must initialize nil to global variable after suit is complete
var _ = AfterSuite(func() {
	CtxTest = nil
	Usr = nil
	UserServerTest = nil
})
