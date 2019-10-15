package main

import (
	"../../pkg/websocket"
	"net/http"
)

func init() {
	websocket.InitClientMap()
}

func main() {
	http.HandleFunc("/ws", websocket.WsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
