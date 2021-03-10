package cleanup

import (
	"alfred/logger"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"io"
)

//more cleanup type will be added here
type CleanupConfig struct {
	//Db *gorm.DB
	ClientConnection     *grpc.ClientConn
	GrpcServerConnection *grpc.Server
	JaegerCloser         io.Closer
}

func GetCleanupConfig(
	ClientConnection *grpc.ClientConn,
	GrpcServerConnection *grpc.Server,
	JaegerCloser io.Closer) *CleanupConfig {
	return &CleanupConfig{
		ClientConnection:     ClientConnection,
		GrpcServerConnection: GrpcServerConnection,
		JaegerCloser:         JaegerCloser,
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
			defer config.JaegerCloser.Close()
			logger.Log.Info("cleanup code successfully executed")
			return nil
		},
	})

}
