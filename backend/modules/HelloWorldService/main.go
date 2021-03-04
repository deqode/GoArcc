package HelloWorldService

import (
	"alfred/modules/HelloWorldService/pb"
)

type HelloService struct {
}

func NewHelloService() pb.HelloWorldServiceServer {
	return HelloService{}
}
