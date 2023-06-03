package middlewares

import (
	"github.com/gin-gonic/gin"
	_jwt "github.com/jeffcail/ginframe/server-common/pkg/jwt"
	"github.com/jeffcail/ginframe/server-common/utils/enum"
	"github.com/jeffcail/ginframe/server-user/global"
	"github.com/jeffcail/ginframe/server-user/models"
	"github.com/spf13/cast"
	"strings"
	"time"
)

// CheckUserAccountIsEnable 检测账号登录状态是否被禁用
func CheckUserAccountIsEnable() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		split := strings.Split(header, " ")
		expire := time.Hour * global.Config.Jwt.Expire
		instance := _jwt.NewJwtInstance(expire, global.Config.Jwt.Secret)

		claims, _ := instance.ParseToken(split[1])
		m := new(models.User)
		u, err := m.FindUserByUsername(claims.Username)
		if err != nil {
			enum.Result.Error(c, enum.ApiCode.FAILED, cast.ToString(err))
			c.Abort()
			return
		}
		if u.Enable == 2 {
			enum.Result.Error(c, enum.ApiCode.USERISDISABLE, enum.ApiCode.GetMessage(enum.ApiCode.USERISDISABLE))
			c.Abort()
			return
		}
		c.Next()
	}
}
