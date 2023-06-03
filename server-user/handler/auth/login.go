package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-common/utils/enum"
	input2 "github.com/jeffcail/ginframe/server-user/input"
	"github.com/jeffcail/ginframe/server-user/service/auth"
	"github.com/spf13/cast"
)

// LoginHandler login struct
type LoginHandler struct{}

// Login admin account login
func (l *LoginHandler) Login(c *gin.Context) {
	input := &input2.LoginInput{}
	if err := c.Bind(input); err != nil {
		enum.Result.Error(c, enum.ApiCode.PARAMERR, cast.ToString(err))
		return
	}
	out, err := auth.LoginService(input)
	if err != nil {
		enum.Result.Error(c, enum.ApiCode.FAILED, cast.ToString(err))
		return
	}

	enum.Result.Success(c, out)
}
