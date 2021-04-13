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