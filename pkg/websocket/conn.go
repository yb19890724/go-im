package websocket

import "sync"

// 映射关系表
var ConnMap *ClientMap

//
type ClientMap struct {
	conn    map[int64]*wsConn
	rwMutex sync.RWMutex
}

func InitClientMap() (clientMap *ClientMap) {
	clientMap = &ClientMap{
		conn: make(map[int64]*wsConn),
	}
	return
}

// 添加链接
func (cl *ClientMap) AddConnection(connId int64, wsConn *wsConn) {

	cl.rwMutex.Lock()
	defer cl.rwMutex.Unlock()

	cl.conn[connId] = wsConn
}

// 剔除链接
func (cl *ClientMap) DelConnection(connId int64) {

	cl.rwMutex.Lock()
	defer cl.rwMutex.Unlock()

	delete(cl.conn, connId)

}

// 获取链接
func (cl *ClientMap) Connection(connId int64) (*wsConn, bool) {

	cl.rwMutex.Lock()
	defer cl.rwMutex.Unlock()

	sConn, ok := cl.conn[connId]

	return sConn, ok
}
