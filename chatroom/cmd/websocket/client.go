package main

import (
	"context"
	"fmt"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close(websocket.StatusInternalError, "内部出错")
	err = wsjson.Write(ctx, conn, "Hello Websocket Server")
	if err != nil {
		panic(err)
	}

	// 读取服务器发来的信息
	var v interface{}
	err = wsjson.Read(ctx, conn,&v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("接收到服务端响应: %v\n", v)
	conn.Close(websocket.StatusNormalClosure, "")
}
