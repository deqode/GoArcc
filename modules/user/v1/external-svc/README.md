# External Service
External Service is isolated from other service. You can write API's which will be exposed as GRPC or REST or GRAPHQL.

## Constructor Calling without grpcClient Connection
```go
func NewUserServer(
	db *gorm.DB,
	config *config.Config,
) pb.UsersServer {
	model.InitialMigrationUser(db)
	return &userServer{
		db:     db,
		config: config,
	}
}
```
## Constructor Calling with grpcClient Connection
```go
type userServer struct {
	db         *gorm.DB
	config     *config.Config
	grpcClient *grpc.ClientConn
}

func NewUserServer(
	db *gorm.DB,
	config *config.Config,
	grpcClientConn *grpc.ClientConn,
) pb.UsersServer {
model.InitialMigrationUser(db)
	return &userServer{
		db:         db,
		config:     config,
		grpcClient: grpcClientConn,
	}
}
```
