package main

import (
	"context"
	"flag"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "wdy",
	})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for i := 0; i < 6; i++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("reps err: %v", resp)
	return nil
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error  {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()

	return nil

}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	// _ = SayHello(client)
	//_ = SayList(client, &pb.HelloRequest{
	//	Name: "wdy",
	//})
	//_ = SayRecord(client, &pb.HelloRequest{
	//	Name: "wdy",
	//})
	_ = SayRoute(client, &pb.HelloRequest{
		Name: "wdy",
	})
}
