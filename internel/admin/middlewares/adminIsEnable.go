package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/internel/admin/models"
	_jwt "github.com/jeffcail/ginframe/pkg/jwt"
	"github.com/jeffcail/ginframe/utils/enum"
	"github.com/spf13/cast"
	"strings"
	"time"
)

// CheckAdminAccountIsEnable 检测账号登录状态是否被禁用
func CheckAdminAccountIsEnable() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		split := strings.Split(header, " ")
		expire := time.Hour * global.Config.Jwt.Expire
		instance := _jwt.NewJwtInstance(expire, global.Config.Jwt.Secret)

		claims, _ := instance.ParseToken(split[1])
		m := new(models.Admin)
		admin, err := m.FindAdminByUsername(claims.Username)
		if err != nil {
			enum.Result.Error(c, enum.ApiCode.FAILED, cast.ToString(err))
			c.Abort()
			return
		}
		if admin.Enable == 2 {
			enum.Result.Error(c, enum.ApiCode.ADMINISDISABLE, enum.ApiCode.GetMessage(enum.ApiCode.ADMINISDISABLE))
			c.Abort()
			return
		}
		c.Next()
	}
}
