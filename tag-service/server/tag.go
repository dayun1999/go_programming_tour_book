package server

import (
	"context"
	"encoding/json"
	"github.com/go-programming-tour-book/tag-service/pkg/bapi"
	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"google.golang.org/grpc/metadata"
	"log"
)

// 【编写gRPC server】 获取标签列表和接口逻辑
type TagServer struct {
	// 3.9节新增auth认证
	auth *Auth
}

type Auth struct {
}

func (a *Auth) GetAppKey() string {
	return "go-programming-book"
}

func (a *Auth) GetAppSecret() string  {
	return "wdywdy"
}

func (a *Auth) Check(ctx context.Context) error  {
	md, _ := metadata.FromIncomingContext(ctx)

	var appKey, appSecret string
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}

	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return errcode.TogRPCError(errcode.Unauthorized)
	}
	return nil
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

// 这就代表了TagServer实现了tag.pb.go文件里面的TagServiceClient接口
func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	//panic("测试抛出异常!")
	if err := t.auth.Check(ctx); err != nil {
		return nil, err
	}
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("md: %+v",md)
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	return &tagList, nil
}
