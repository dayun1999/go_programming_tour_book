package main

import (
	"context"
	"flag"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

var port string

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "你好,这里是服务端...",
	}, nil
}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for i := 0; i < 6; i++ {
		_ = stream.Send(&pb.HelloReply{Message: "[服务端流式RPC] 你好,stream"})
	}
	return nil
}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{
				Message: "[客户端流式RPC] say.Record",
			}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v", resp)
	}
}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{
			Message: "say.route",
		})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)

		n++
	}
}

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	_ = server.Serve(lis)
}
