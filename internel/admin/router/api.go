package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/internel/admin/handler"
	"github.com/jeffcail/ginframe/internel/admin/handler/auth"
)

// ApiRouter http api routers
func ApiRouter(r *gin.Engine) {

	r.GET("/ping", handler.Ping)
	r.GET("/list", handler.PagePagination)

	v1Group := r.Group("/api/v1")
	setAuthRouter(v1Group)

}

// 登录授权
func setAuthRouter(v1Group *gin.RouterGroup) {
	loginHandler := new(auth.LoginHandler)
	v1Group.POST("/login", loginHandler.Login)
}
