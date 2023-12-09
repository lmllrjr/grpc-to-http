package server

import (
	"context"

	"github.com/golang/glog"
	"github.com/lmllrjr/grpc-to-http/pb"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Implements of pb.GreeterServer
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func newGreeterServer() pb.GreeterServer {
	return new(greeterServer)
}

func (s *greeterServer) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetReply, error) {
	glog.Info("method: Greet", in)
	return &pb.GreetReply{Message: "Hello " + in.Name}, nil
}

func (s *greeterServer) HelloWorld(ctx context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	glog.Info("method: HelloWorld")
	return &httpbody.HttpBody{
		ContentType: "application/json",
		Data:        []byte(`{"message":"hello world"`),
	}, nil
}
