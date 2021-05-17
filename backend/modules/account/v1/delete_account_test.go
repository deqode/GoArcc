package account_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Describe:Delete account test", func() {
	/*	var (
			accountServer     pb.AccountsServer
		)
		//now init logger
		logger.Init(&logger.Config{
			LogLevel:    zap.DebugLevel, // TODO: Take this level from config
			Development: false,
		})
		//getting config
		cfgFile, err := config.LoadConfig("config", "/home/deq/Desktop/deqodeEnviroment/platform/backend")
		if err != nil {
			log.Fatal(err)
		}
		cfg, err := config.ParseConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		fields :=  struct {
			db         *gorm.DB
			config     *config.Config
			grpcClient *grpc.ClientConn

		}{
			db:                db.NewConnection(cfg),
			config:            cfg,
			grpcClient:        grpcClient.GetGrpcClientConnection(cfg),
		}
		//service initialisation
		accountServer = account.NewAccountsServer(fields.db,fields.config,fields.grpcClient)
		Describe("Describe:Delete account when id is incorrect", func() {
			Context("Context:Account with incorrect id", func() {
				It("It: Delete with Incorrect id", func() {
					_ , err := accountServer.DeleteAccount(context.Background() , &pb.DeleteAccountRequest{
						Id: "",
					})
					Expect(err).Should(Equal(status.Error(codes.NotFound , "record not found")))
				})
			})
		})

		Describe("Describe:When Request object is nil", func() {
			Context("Context:Context:When Request Object is nil", func() {
				It("It:When Request Object is nil", func() {
					_ , err := accountServer.DeleteAccount(context.Background() , nil)
					Expect(err).Should(Equal(status.Error(codes.FailedPrecondition , "request object must be present")))
				})
			})
		})
	*/
})
