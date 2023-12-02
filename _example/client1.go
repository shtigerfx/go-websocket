package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func loop_send(conn *websocket.Conn) {
	// 循环发送消息
	for {
		// 发送消息
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, Galori!"))
		if err != nil {
			fmt.Println("发送消息失败：", err)
			return
		}

		// 等待 5 秒钟
		time.Sleep(5 * time.Second)
	}
}

func loop_receive(conn *websocket.Conn) {
	// 循环接收消息
	for {
		// 接收消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("接收消息失败：", err)
			return
		}
		fmt.Println("收到消息：", string(message))

		// 等待 5 秒钟
		time.Sleep(5 * time.Second)
	}
}

func main() {
	// // 定义 WebSocket 协议版本
	// var upgrader = websocket.Upgrader{
	// 	ReadBufferSize:  1024,
	// 	WriteBufferSize: 1024,
	// 	CheckOrigin: func(r *http.Request) bool {
	// 		return true
	// 	},
	// }

	// 连接 WebSocket 服务器
	// conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:7800/ws?systemId=live", nil)
	conn, _, err := websocket.DefaultDialer.Dial("ws://43.129.188.163:7800/ws?systemId=live", nil)
	if err != nil {
		fmt.Println("连接 WebSocket 服务器失败：", err)
		return
	}
	defer conn.Close()

	go loop_receive(conn)

	go loop_send(conn)

	time.Sleep(5 * time.Hour)
}
