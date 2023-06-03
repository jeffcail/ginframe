package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-user/global"
	"github.com/jeffcail/ginframe/server-user/router"
)

// WebsocketServer websocket服务
func WebsocketServer() error {
	r := gin.Default()
	router.WsRouter(r)
	return r.Run(fmt.Sprintf("%s%s", ":", global.Config.Websocket.BindPort))
}
