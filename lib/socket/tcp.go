package socket

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	isClosed  bool
	mutex     sync.Mutex
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	//启动读协程
	go conn.readLoop()

	//启动写协程
	go conn.writeLoop()

	return
}

//API
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		// fmt.Println("connection is closed")
	}
	//fmt.Println(conn.wsConn.RemoteAddr().String())
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	//fmt.Println(string(data[:]))
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		// fmt.Println("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	// 线程安全的close，可重入
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}

		//阻塞在这里，等待inChan有空位置
		//但是如果writeLoop连接关闭了，这边无法得知
		//conn.inChan <- data

		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			//closeChan关闭的时候，会进入此分支
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR

		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
		// conn.outChan <- data
	}
ERR:
	conn.Close()
}
