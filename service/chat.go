package service

import (
	"../lib/socket"
	"github.com/gorilla/websocket"
	"net/http"
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
	//fmt.Println(r)
	if conn, err = socket.InitConnection(wsConn); err != nil {
		goto ERR
	}
	// 服务端保持长连接，一般由前端发送消息保持长连接亦可
	//go func() {
	//	var (
	//		err error
	//	)
	//	for {
	//		// 5s发一次心跳包，保持连接状态
	//		if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
	//			return
	//		}
	//		time.Sleep(5 * time.Second)
	//	}
	//}()

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

func InitChat(port string) {
	http.HandleFunc("/", wsHandler)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}