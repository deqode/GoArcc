package graphql

import (
	"alfred/logger"
	accPb "alfred/modules/account/v1/pb"
	architectureSuggesterPb "alfred/modules/architecture-suggester/pb"
	gitPb "alfred/modules/git/v1/pb"
	userProfilePb "alfred/modules/user-profile/v1/pb"
	vcsPb "alfred/modules/vcs-connection/v1/pb"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterGraphqlModules Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGraphqlModules: Mapping the services with the single graphql endpoint
*/
func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := userProfilePb.RegisterUserProfilesGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := gitPb.RegisterGitsGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := architectureSuggesterPb.RegisterArchitectureSuggesterServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := vcsPb.RegisterVCSConnectionsGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := accPb.RegisterAccountsGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
