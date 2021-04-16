package main

import (
	"context"
	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

// 编写示例来调用gRPC服务
func main() {
	var opts []grpc.DialOption
	auth := Auth{
		AppKey: "go-programming-book",
		AppSecret: "wdywdy",
	}
	ctx := context.Background()
	newCtx := metadata.AppendToOutgoingContext(ctx,"wdywdy", "go语言") // metadata, 客户端
	opts = append(opts, grpc.WithPerRPCCredentials(&auth))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(),
			middleware.ClientTracing(),
			)))

	clientCon, _ := GetClientConn(newCtx, "localhost:8004", opts) // 加上WithBlock()会阻塞直到连接建立
	defer clientCon.Close()
	targetServiceClient := pb.NewTagServiceClient(clientCon)
	resp, _ := targetServiceClient.GetTagList(newCtx, &pb.GetTagListRequest{Name: "Go"})
	log.Printf("resp: %v", resp)
}


func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...) // DialContext()并不能立刻创建连接,因为其是非阻塞的
}

// 实现Auth认证
type Auth struct {
	AppKey string
	AppSecret string
}

// 实现grpc里面的默认认证接口PerRPCCredentials
func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
	return map[string]string{
		"app_key": a.AppKey,
		"app_secret": a.AppSecret,
	}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

