package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"baishi/grpc-demo/proto/demo"
)

func NewDemoServer() demo.DemoServer {
	return &demoServer{}
}

type demoServer struct {
	demo.UnsafeDemoServer
}

func (d *demoServer) HelloWorld(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
	return nil, nil
}
