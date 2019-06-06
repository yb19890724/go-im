package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)


var rwLock  sync.RWMutex// 读写锁

// 客户端读写消息
type wsMessage struct {
	messageType int
	data        []byte
}

type Message struct {
	Token   string  `json:"token"`
	Context string  `json:"context"`
}

// 客户端连接
type wsConn struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列
	uid      int64
	token    string
	
	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

// 与客户端建立web socket 连接
func upgrade(w http.ResponseWriter, r *http.Request) (conn *websocket.Conn, err error) {
	
	conn,err =(&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w,r,nil)
	
	
	// 应答客户端告知升级连接为websocket
	return conn,err
}

// ws 调度
func WsHandler(w http.ResponseWriter, r *http.Request) {
	
	
	r.ParseForm()
	
	token :=r.Form.Get("token")
	
	// todo:  验证用户auth
	uid,err:=checkAuth(token) // 校验token
	
	if err != nil {
		return
	}
	
	conn, err := upgrade(w, r)
	
	if err != nil {
		return
	}
	
	fmt.Println(fmt.Sprintf("client %d connection server success",uid))
	
	wsConn := &wsConn{
		wsSocket:  conn,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		uid:       uid,
		token:     token,
		closeChan: make(chan byte),
		isClosed:  false,
	}
	
	ConnMap.AddConnection(uid,wsConn)
	
	// 处理器
	go procLoop(wsConn)
	
	go wsReadLoop(wsConn)
	go wsResponse(wsConn)
	go wsWriteLoop(wsConn)
	
	
}


// 读协程
func wsReadLoop(wsConn *wsConn)  {
	
	for {
		// 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		
		req := &wsMessage{
			msgType,
			data,
		}
		
		// 放入请求队列
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}


// 发送协程
func wsWriteLoop(wsConn *wsConn) {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			
			var sendMsg Message
			
			err :=json.Unmarshal([]byte(msg.data), &sendMsg)
			
			if err != nil{
				fmt.Println(err)
			}
			
			uid, err := checkAuth(sendMsg.Token)
			
			fmt.Println(fmt.Sprintf("send client user:%s",uid))
			
			sConn,ok := ConnMap.Connection(uid)
			
			if ok {
				if err := sConn.wsSocket.WriteMessage(msg.messageType, []byte(sendMsg.Context)); err != nil {
					sConn.wsClose()
				}
			}
			// 写给websocket
			
		case <-wsConn.closeChan:
			goto closed
		}
	}
closed:
}

// 读通道写入到写通道
func wsResponse(wsConn *wsConn) {
	for v := range wsConn.inChan {
		wsConn.outChan <- v
	}
}


// 退出关闭所有资源
func (wsConn *wsConn) wsClose() {
	// 关闭当前长连接
	wsConn.wsSocket.Close()
	
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan) // 关闭通知管道，通知所有协程退出
	}
	
	if _, ok := ConnMap.Connection(wsConn.uid); ok {// 关闭资源删除用户连接
		// 删除用户连接记录
		
		ConnMap.DelConnection(wsConn.uid)
		fmt.Println(fmt.Sprintf("clear client:%d socket",wsConn.uid))
	}
}


// 心跳检测
func procLoop(wsConn *wsConn) {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			
			msg, _ := json.Marshal(Message{
				Token:   wsConn.token,
				Context: "heartbeat from server",
			})
			
			if err := wsConn.wsWrite(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.wsClose()
				break
			}
		}
	}()
}

func (wsConn *wsConn)wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data,}:
	case <- wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

// 读取消息
func (wsConn *wsConn)wsRead() (*wsMessage, error) {
	select {
	case msg := <- wsConn.inChan:
		return msg, nil
	case <- wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}





// 验证token
func checkAuth(token string) (uid int64, err error) {
	
	if token == "11111" {
		return 11111, nil
	}
	
	if token == "22222" {
		return 22222, nil
	}
	
	return 0, errors.New("user auth error")
}