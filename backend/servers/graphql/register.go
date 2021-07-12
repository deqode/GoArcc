package graphql

import (
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"goarcc/logger"
	userProfilePb "goarcc/modules/user-profile/v1/pb"
	"google.golang.org/grpc"
)

// RegisterGraphqlModules Todo : Whenever any new modules will be in goarcc : it must be registered in below method
/*
  RegisterGraphqlModules: Mapping the services with the single graphql endpoint
*/
func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := userProfilePb.RegisterUserProfilesGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
