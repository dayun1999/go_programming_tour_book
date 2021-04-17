package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 客户端首先去连接服务器
	conn, err := net.Dial("tcp", ":2020")
	if err != nil {
		panic(err)
	}
	done := make(chan struct{})
	go func() {
		// 将控制台的输入发送到连接里面
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		// 通知goroutine退出
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _,err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
