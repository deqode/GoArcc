package helpers

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"google.golang.org/grpc"
)

func NewClient() *grpc.ClientConn {
	c := config.Config{Grpc: config.GraphqlServerConfig{
		Port: "8080",
		Host: "localhost",
	}}
	clientCon := grpcClient.GetGrpcClientConnection(&c)
	return clientCon
}
