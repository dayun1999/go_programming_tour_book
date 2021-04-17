# go_programming_tour_book

《Go编程之旅 一起用go做项目》<br>
首先要设置国内镜像代理:
```go
go env -w GORPOXY=https://goproxy.cn,direct
```

### 项目里用到的第三方库

- #### P116 邮件报警处理
```go
go get -u gopkg.in/gomail.v2
```
- #### P121 接口限流控制
```go
go get -u github.com/juju/ratelimit
```

- #### P128 链路追踪
```go
go get -u github.com/opentracing/opentracing-go
go get -u github.com/uber/jaeger-client-go
```

- #### P135 SQL追踪
```go
go get -u github.com/eddycjy/opentracing-gorm
```

- #### P141 配置热更新
```go
go get -u golang.org/x/sys/...
go get -u github.com/fsnotify/fsnotify
```

- #### P164 安装Protobuf编译器
```shell script
wget <对应网址>
unzip <protobuf>.zip && cd <解压后的文件>
./autogen.sh
./configure
make
make check
make install
// (环境为CentOS 7)下载好对应的版本(3.15)之后, 发现make失败
// 如何解决?
// 安装autoconf和automake
yum -y install gcc automake autoconf libtool make

// 安装g++:
yum install gcc gcc-c++
```

- #### P165 安装protoc插件
```go
go get -u github.com/golang/protobuf/protoc-gen-go
``` 

- #### P171 安装gRPC
```go
go get -u google.golang.org/grpc
```

- #### P189 调试gRPC接口
```go
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

- #### P202 在同端口监听HTTP
```go
go get -u github.com/soheilhy/cmux
```

- #### P204 同端口同方法提供双流量支持
```go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
// 将编译的Proto Plugin的可执行文件从Gopath中移到go安装目录下
mv $GOPATH/bin/protoc-gen-grpc-gateway /usr/local/go/bin/

// 对proto文件的处理,-I参数的作用是指定import搜索的目录,见P205
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto

```

- #### P213 接口文档
```go
// 安装插件
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
// 安装swagger-ui
// 下载 源码包zip，解压之后将dist目录下的所有文件拷贝到third_party/swagger-ui/下
```

- #### P213 静态资源转换
```go
// 将资源文件转换为go代码
go get -u github.com/go-bindata/go-bindata/...
```

- #### P213 Swagger UI 的处理和访问
```go
go get -u github.com/elazarl/go-bindata-assetfs/...
```

- #### P215 Swagger 描述文件的生成和读取
```go
// 生成swagger定义文件
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --swagger_out=logtostderr=true:. ./proto/*.proto
```

- #### P220 用多个拦截器
```go
go get -u github.com/grpc-ecosystem/go-grpc-middleware
```

- #### P236 链路追踪(gRPC + jaeger)
```go
go get -u github.com/opentracing/opentracing-go
go get -u github.com/uber/jaeger-client-go
```

- #### P281 Websocket的使用
```go
go get nhooyr.io/websocket
```
wireshark抓包协议分析<br>
<p align="center">
    <img src="https://github.com/code4EE/go_programming_tour_book/blob/main/20210417181104.png" width="600">
</p>

