# go_programming_tour_book
### 2025年06月27日 修改 <br>
今天无意间发现这个4年前的仓库有了10几个star, 我脑子里面一闪而过的就是4年前那个迷茫学习编程的自己, 如今我已经工作3年了, 趁着有点时间分享一下自己的一些浅薄的经验吧。<br>

#### 语言并不重要
编程语言确实不重要(但是起码要熟悉一门语言并且作为自己的主力语言), 基础知识和学习能力才是通用的, 基础知识包括但不限于数据结构、基础算法、网络知识、计算机组成(内存、CPU、缓存、磁盘等)、项目架构分层(经典的就MVC那套)、各个语言的特点(比如 Golang 到底适合什么场景);<br>
这里顺便补充一些个人的见闻: <br>
目前国内这一批使用 Golang 的大中小企业中, Golang的大部分使用场景如下: 容器(k8s相关)、中间件、Agent代理、网络编程(v2ray-core)、APM领域(监控、日志、链路追踪)、命令行工具等<br>

#### 软素质
其实编程能力和项目知识都属于硬技能, 也是面试里面能直接看到的东西。但是软素质在工作当中也很重要, 包括方案设计/调研能力、问题分析&解决能力、任务拆分和排优先级能力、沟通能力、总结能力等, 总之新入职的同学要记住: 做任何事情不要闷头做, 看不懂的方案和问题及时问清楚, 不然容易出现代码返工的情况, 顺便需要了解一下第一性原理。

#### 如何写好代码
学习写优秀代码的过程都涉及到所谓的"借鉴", 刚刚开始工作的时候, 其实就是看你能不能快速的"抄"到答案, 这个"抄"不是贬义, 有些代码就需要那么写, 多看多写两遍就能学会了。但是工作时间一拉长, 会发现有大部分时间都是在写面条代码, 也就是网上很多人吐槽的工作了几年都是在增删改查, 其实本质上编程能力压根没有进步。<br>
所以在不同阶段需要干不同的事情————
1. 新接触业务代码的时候, 一定要搞清楚项目是怎么运行的, 都涉及到哪些模块, 比如鉴权、数据存储等是在哪里做的, 自己按照"数据流向"的方式梳理清楚(比如你的后端服务是如何完整处理请求的, 都经过哪些模块?) 
2. 工作了一段时间之后, 这时候就需要积累自己的编程知识了, 可以慢慢看一些好的甚至杂七杂八的网站, 也可以直接看优秀的、篇幅不大的项目进行学习(比如ants、zap等仓库的源码)；也可以看下一些经典的设计模式的应用([go-patterns](https://github.com/dayun1999/go-patterns)), 比如options模式、代理模式、Fan-in/Fan-out等
   
#### 关于业务和基础架构
我个人把写代码的工作分为了2类:
1. 一种叫基础设施开发, 比如像网关、存储、监控、告警、链路追踪等等这些, 基本上编码能力方面的要求比较高, 因为会涉及到高性能, 基础设施就是不应该容易出错。
2. 一种叫应用产品开发, 对于大部分公司来说, 招聘人员是为了搞定产品, 搞定产品是为了卖钱, 所以需要分清楚自己所在的公司到底做的是什么产品, 更关键的是自己现在在干什么, 自己将来想干什么, 从而自己规划自己的职业路线学习对应的知识。 

#### 关于AI
目前基本上人人都能用到大模型, ChatGPT、DeepSeek、豆包等在工作中确实能提升工作效率, 很多代码它们都能帮你写出来, 从而帮助我们从编码细节解脱出来, 当然目前来说提问的篇幅是个问题, 这个会考验你个人的任务&功能拆分, 起码当下是这样的<br>
所以其实提到AI的时候, 我觉得对待编程的学习, 可以暂时抛弃掉之前的那种方式了（之前是通过慢慢的阅读书籍、论文、各种网站找面试题强化自己）, 现在可以直接引导AI帮助自己学习, 不懂的方向和细枝末节大模型现在完全有能力帮助你。<br>
另外关于AI会不会替代程序员, 我目前也没有一个对未来3-5年的场景的猜测, 不过长期趋势是这样的。

### 2021年09月06日 修改
### 【PS】新来的初学者别看我这个了，直接去看原作者之一(煎鱼大佬)开源的图书[点这里](https://golang2.eddycjy.com/)
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
![image](https://github.com/code4EE/images/blob/main/20210417181104.png)

