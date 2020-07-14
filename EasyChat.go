package main

import (
	"./lib/socket"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *socket.Connection
		data   []byte
	)

	//Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	fmt.Println(r.RequestURI)
	if conn, err = socket.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	//关闭连接
	conn.Close()
}

func main() {
	//http:localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
