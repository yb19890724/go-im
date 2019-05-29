package main

import (
	"github.com/yb19890724/go-im/pkg/websocket"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", websocket.WsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
