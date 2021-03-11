package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
)

type HelloService struct {
}

//Service Implementation
func NewHelloService() pb.HelloWorldServiceServer {
	return HelloService{}
}

func LocalHelloWorld() pb.HelloWorldServiceServer {
	return HelloService{}
}
