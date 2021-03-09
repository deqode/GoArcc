package cleanup

import (
	"alfred/logger"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

//more cleanup type will be added here
type CleanupConfig struct {
	//Db *gorm.DB
	ClientConnection     *grpc.ClientConn
	GrpcServerConnection *grpc.Server
}

func GetCleanupConfig(
	ClientConnection *grpc.ClientConn,
	GrpcServerConnection *grpc.Server) *CleanupConfig {
	return &CleanupConfig{
		ClientConnection:     ClientConnection,
		GrpcServerConnection: GrpcServerConnection,
	}
}

func Cleanup(lc fx.Lifecycle, config *CleanupConfig) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Log.Info(".......Starting Cleanup code ......")
			//closing db connection
			if err := config.ClientConnection.Close(); err != nil {
				logger.Log.Info("not able to close the grpc client connection")
			}
			//Closing grpc server connection
			config.GrpcServerConnection.GracefulStop()
			logger.Log.Info("successfully closed grpc server connection")
			//todo more cleanup code will be added
			logger.Log.Info("cleanup code successfully executed")
			return nil
		},
	})

}
