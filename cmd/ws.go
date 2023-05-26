package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/internel/router"
)

// WebsocketServer websocket服务
func WebsocketServer() error {
	r := gin.Default()
	router.WsRouter(r)
	return r.Run(fmt.Sprintf("%s%s", ":", global.Config.Websocket.BindPort))
}
