package grpcClient_test

import (
	"alfred/client/grpcClient"
	"alfred/config"
	"google.golang.org/grpc"
	"testing"
)

//Unit Test covering all the functionality test
//Note : It will not cover , API Testing or bdd Test cases
//Logical part of code will be tested
func TestGetGrpcClientConnectionWithNoGrpcPort(t *testing.T) {
	//must return empty method config , because no server started
	actualClientConnection := grpcClient.GetGrpcClientConnection(&config.Config{GRPCPort: ""}, grpcClient.GetGrpcClientConnectionContext())
	actualMethodConfig := actualClientConnection.GetMethodConfig("abc/deg")
	var expectedMethodConfig grpc.MethodConfig
	if actualMethodConfig != expectedMethodConfig {
		t.Fatal("expected config does not match with the actual config", "actualConfig: ",
			actualMethodConfig, " expectedConfig: ", expectedMethodConfig)
	}
}
