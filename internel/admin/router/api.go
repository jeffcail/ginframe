package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/internel/admin/handler"
)

// ApiRouter http api routers
func ApiRouter(r *gin.Engine) {

	r.GET("/ping", handler.Ping)
	r.GET("/list", handler.PagePagination)

}
