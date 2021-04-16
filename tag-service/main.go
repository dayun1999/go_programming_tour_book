package main

import (
	"context"
	"encoding/json"
	"flag"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	"github.com/go-programming-tour-book/tag-service/pkg/swagger"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"path"
	"strings"
)

var (
	// 双端口各自监听http和rpc协议
	//grpcPort string
	//httpPort string

	// 在同端口监听HTTP
	port string
)

// 自定义错误
type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message, omitempty"`
}

func init() {
	//flag.StringVar(&httpPort, "http_port", "9001", "HTTP启动端口号")
	//flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC启动端口号")

	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func main() {
	// 版本3
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
	// 版本2
	//l, err := RunTCPServer(port)
	//if err != nil {
	//	log.Fatalf("Run TCP Server err: %v", err)
	//}
	//m := cmux.New(l)
	//grpcL := m.MatchWithWriters(
	//	cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
	//)
	//httpL := m.Match(cmux.HTTP1Fast())
	//
	//grpcS := RunGrpcServer()
	//httpS := RunHttpServer(port)
	//go grpcS.Serve(grpcL)
	//go httpS.Serve(httpL)
	//
	//err = m.Serve()
	//if err != nil {
	//	log.Fatalf("Run Server err: %v", err)
	//}

	// 版本1
	//errs := make(chan error)
	//// 为什么要启动两个goroutine？因为RunHttpServer和RunGrpcServer都是阻塞行为
	//go func() {
	//	err := RunHttpServer(httpPort)
	//	if err != nil {
	//		errs <- err
	//	}
	//}()
	//go func() {
	//	err := RunGrpcServer(grpcPort)
	//	if err != nil {
	//		errs <- err
	//	}
	//}()
	//
	//select {
	//case err := <-errs:
	//	log.Fatalf("Run Server err: %v", err)
	//}

	// 下面的代码已经被封装起来了,见函数RunGrpcServer
	//s := grpc.NewServer()
	//pb.RegisterTagServiceServer(s, server.NewTagServer())
	//reflection.Register(s)
	//
	//lis, err := net.Listen("tcp", ":"+port)
	//if err != nil {
	//	log.Fatalf("net.Listen err: %v", err)
	//}
	//
	//err = s.Serve(lis)
	//if err != nil {
	//	log.Fatalf("server.Serve err: %v", err)
	//}
}

// 版本3----------------------同端口同方法提供双流量支持------------------------------
// 不同协议的分流
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	//
	//endpoint := "0.0.0.0:" + port
	//runtime.HTTPError = grpcGatewayError
	//gwmux := runtime.NewServeMux()
	gatewayMux := runGrpcGatewayServer()
	httpMux.Handle("/", gatewayMux)
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong\n"))
	})

	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset: swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix: "third_party/swagger-ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)
		http.ServeFile(w, r, p)
	})
	return serveMux
}

func runGrpcServer() *grpc.Server {
	// 实现一个 / 多个拦截器
	// gRPC Server的相关属性的配置可以在写这
	opts := []grpc.ServerOption {
		//grpc.UnaryInterceptor(HelloInterceptor),
		//grpc.UnaryInterceptor(WorldInterceptor),
		// 上述会报错: panic: The unary server interceptor was already set and may not be reset.

		// 使用第三方库go-grpc-middleware,链式方式达到用多个拦截器的目的
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			HelloInterceptor,
			WorldInterceptor,
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
			middleware.ServerTracing,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}
// 一元拦截器之一
func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)  (interface{}, error) {
	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")
	return resp, err
}
// 一元拦截器之一
func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)  (interface{}, error) {
	log.Println("你好, 世界")
	resp, err := handler(ctx, req)
	log.Println("再见, 世界")
	return resp, err
}
func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts) // 注册TagServiceHandler事件
	return gwmux
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{
		Code:    int32(s.Code()),
		Message: s.Message(),
	}

	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-Type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))

	_, _ = w.Write(resp)
}

// 版本2------------------同端口监听HTTP------------------------------------
// 针对HTTP协议
//func RunHttpServer(port string) *http.Server {
//	serverMux := http.NewServeMux() // 初始化一个多路复用器
//	serverMux.HandleFunc("/ping",
//		func(w http.ResponseWriter, r *http.Request) {
//			_, _ = w.Write([]byte(`pong`))
//		},
//	)
//	return &http.Server{
//		Addr:    ":" + port,
//		Handler: serverMux,
//	}
//}

// 针对gRPC协议
//func RunGrpcServer() *grpc.Server {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s) // 使用工具grocurl的前提是gRPC server已经注册了反射服务
//	return s
//}
//
//// 针对TCP协议
//func RunTCPServer(port string) (net.Listener, error) {
//	return net.Listen("tcp", ":"+port)
//}

// 版本1-----------------------另起端口监听------------------------------

// 将原本的gRPC服务启动端口调整为对HTTP1.1 和 gRPC端口号的读取
// 针对HTTP的RunHttpServer方法,可用于基本的心跳检测
//func RunHttpServer(port string) error {
//	serverMux := http.NewServeMux() // 初始化一个多路复用器
//	serverMux.HandleFunc("/ping",
//		func(w http.ResponseWriter, r *http.Request) {
//			_, _ = w.Write([]byte(`pong`))
//		},
//	)
//	return http.ListenAndServe(":"+port, serverMux)
//}

// 针对gRPC--版本1
//func RunGrpcServer(port string) error {
//	s := grpc.NewServer()
//	pb.RegisterTagServiceServer(s, server.NewTagServer())
//	reflection.Register(s) // 使用工具grocurl的前提是gRPC server已经注册了反射服务
//	lis, err := net.Listen("tcp", ":"+port)
//	if err != nil {
//		return err
//	}
//	return s.Serve(lis)
//}
