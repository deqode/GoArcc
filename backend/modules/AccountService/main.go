package AccountService

//type AccountService struct {
//	db         *gorm.DB
//	config     *config.Config
//	grpcClient *grpc.ClientConn
//}
//
////todo : AlWays add migration code for best practices
//func NewAccountService(db *gorm.DB, config *config.Config, grpcClientConn *grpc.ClientConn) pb.AccountServiceServer {
//
//	//initial migration of databases: schema migration
//	models.InitialMigrationAccount(db)
//	return &AccountService{
//		db:         db,
//		config:     config,
//		grpcClient: grpcClientConn,
//	}
//}
