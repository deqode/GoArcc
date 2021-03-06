package cleanup

import (
	"context"
	"github.com/deqode/GoArcc/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"io"
)

// Config CleanupConfig : cleanup config is the configuration of closing the instance.
// more type will be added in cleanup config if we stop or destroy the instance.
type Config struct {
	//Db *gorm.DB
	GrpcServerConnection *grpc.Server
	GrpcClientConnection *grpc.ClientConn
	JaegerCloser         io.Closer
}

// GetCleanupConfig GetCleanupConfig: Get cleanup config is the constructor.
//  required all the closing instance
func GetCleanupConfig(
	GrpcServerConnection *grpc.Server,
	GrpcClientConnection *grpc.ClientConn,
	JaegerCloser io.Closer) *Config {
	return &Config{
		GrpcServerConnection: GrpcServerConnection,
		JaegerCloser:         JaegerCloser,
		GrpcClientConnection: GrpcClientConnection,
	}
}

// Cleanup /*
func Cleanup(lc fx.Lifecycle, config *Config) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Log.Info(".......Starting Cleanup code ......")

			//Closing client connection
			config.GrpcClientConnection.Close()
			logger.Log.Info("successfully closed grpc client connection")
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
