package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/internel/admin/handler"
)

// WsRouter websocket routers
func WsRouter(r *gin.Engine) {
	r.GET("/ws", handler.WsPing)
}
