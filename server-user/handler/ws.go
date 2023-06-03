package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jeffcail/ginframe/server-user/ws"
	"log"
	"time"
)

var (
	//websocket 长连接
	wsConn *websocket.Conn
	err    error
	conn   *ws.Connection
	data   []byte
)

// WsPing ws router test
func WsPing(c *gin.Context) {

	wsConn, err = ws.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("http upgrade tcp failed", err)
	}

	conn, err = ws.InitConnection(wsConn)
	if err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.WriteMessage([]byte("hello ginframe")); err != nil {
				return
			}
			time.Sleep(time.Second * 1)
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
	conn.Close()
}
