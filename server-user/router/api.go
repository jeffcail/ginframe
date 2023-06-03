package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-user/handler/auth"
	"github.com/jeffcail/ginframe/server-user/handler/user"
	"github.com/jeffcail/ginframe/server-user/middlewares"
)

// ApiRouter http api routers
func ApiRouter(r *gin.Engine) {

	v1Group := r.Group("/api/v1")
	setAuthRouter(v1Group)
	setUserRouter(v1Group)

}

// 登录授权
func setAuthRouter(v1Group *gin.RouterGroup) {
	loginHandler := new(auth.LoginHandler)
	v1Group.POST("/login", loginHandler.Login)
}

// 用户
func setUserRouter(v1Group *gin.RouterGroup) {
	userHandler := new(user.UserHandler)
	user := v1Group.Group("/user")
	user.Use(middlewares.AuthMiddleware())
	user.Use(middlewares.CheckUserAccountIsEnable())
	user.GET("/list", userHandler.List)
}
