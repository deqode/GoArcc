package HelloWorldService

import (
	"alfred/logger"
	"alfred/modules/HelloWorldService/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h HelloService) HelloWorld(ctx context.Context, hello *pb.Hello) (*empty.Empty, error) {
	logger.Log.Info("Hello")
	panic("just checking panic")
	return &empty.Empty{}, nil
}
