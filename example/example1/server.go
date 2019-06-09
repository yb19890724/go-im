package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列
	Exit     context.CancelFunc
}

func initConn(ws *websocket.Conn, cancel context.CancelFunc) *wsConnection {
	
	return &wsConnection{
		wsSocket: ws,
		inChan:   make(chan *wsMessage, 1000),
		outChan:  make(chan *wsMessage, 1000),
		Exit:     cancel,
	}
}

func (wsConn *wsConnection) wsReadLoop(ctx context.Context) {
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto ERROR
		}
		fmt.Println(string(data))
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-ctx.Done():
			goto CLOSED
		}
	}
ERROR:
	wsConn.Exit()
CLOSED:
	fmt.Println("wsReadLoop退出")
}

func (wsConn *wsConnection) wsWriteLoop(ctx context.Context) {
	
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto ERROR
			}
		case <-ctx.Done():
			goto CLOSED
		}
	}
ERROR:
	wsConn.Exit()
CLOSED:
	fmt.Println("wsWriteLoop退出")
}

func (wsConn *wsConnection) processLoop(ctx context.Context) {
	for {
		select {
		case msg := <-wsConn.inChan:
			wsConn.outChan <- msg
		case <-ctx.Done():
			goto CLOSED
		}
	}
CLOSED:
	fmt.Println("processLoop退出")
}

func (wsConn *wsConnection) heartBeatLoop(ctx context.Context) {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.wsWrite(ctx, websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.Exit()
				break
			}
		}
	}()
}

func wsHandler(resp http.ResponseWriter, req *http.Request) {
	
	ctx, cancel := context.WithCancel(context.Background())
	
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	
	if err != nil {
		return
	}
	
	fmt.Println("connection success")
	
	wsConn := initConn(wsSocket, cancel)
	
	// 心跳
	go wsConn.heartBeatLoop(ctx)
	// 处理器
	go wsConn.processLoop(ctx)
	// 读协程
	go wsConn.wsReadLoop(ctx)
	// 写协程
	go wsConn.wsWriteLoop(ctx)
}

func (wsConn *wsConnection) wsWrite(ctx context.Context, messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data,}:
	case <-ctx.Done():
		return errors.New("websocket closed")
	}
	return nil
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
