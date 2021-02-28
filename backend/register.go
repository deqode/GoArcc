package main

import (
	"alfred/UserProfile"
	userprofilePb "alfred/UserProfile/pb"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

//TODO : New Modules will be registered here
func RegisterModules(server *grpc.Server)  {
	//Registering User profile Modules
	userprofilePb.RegisterUserProfileServiceServer(server, &UserProfile.UserServer{})

}



//Todo : GraphQl Module will be registered here
func RegisterGraphqlModules(runtime *runtime.ServeMux)  error {
	if err := userprofilePb.RegisterUserProfileServiceGraphql(runtime) ; err != nil {
		return err
	}
	return nil
}
