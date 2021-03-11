package cleanup

import (
	"alfred/logger"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"io"
)

/*
  CleanupConfig: cleanup config is the configuration of closing the instance.
  more type will be added in cleanup config if we stop or destroy the instance.
*/
type CleanupConfig struct {
	//Db *gorm.DB
	GrpcServerConnection *grpc.Server
	JaegerCloser         io.Closer
}

/*
  GetCleanupConfig: Get cleanup config is the constructor.
  required all the closing instance
*/
func GetCleanupConfig(
	GrpcServerConnection *grpc.Server,
	JaegerCloser io.Closer) *CleanupConfig {
	return &CleanupConfig{
		GrpcServerConnection: GrpcServerConnection,
		JaegerCloser:         JaegerCloser,
	}
}

/*
  Cleanup: cleanup is the lifecycle app .
  responsible for destroying and closing the
  instance.
*/
func Cleanup(lc fx.Lifecycle, config *CleanupConfig) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Log.Info(".......Starting Cleanup code ......")
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
