package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
	_jwt "github.com/jeffcail/ginframe/pkg/jwt"
	"github.com/jeffcail/ginframe/utils/enum"
	"strings"
	"time"
)

// AuthMiddleware auth token check
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			enum.Result.Error(c, enum.ApiCode.TOKENISNOTEXISTS, enum.ApiCode.GetMessage(enum.ApiCode.TOKENISNOTEXISTS))
			c.Abort()
			return
		}

		split := strings.Split(header, " ")
		if len(split) != 2 {
			enum.Result.Error(c, enum.ApiCode.TOKENISVALID, enum.ApiCode.GetMessage(enum.ApiCode.TOKENISVALID))
			c.Abort()
			return
		}
		expire := time.Hour * global.Config.Jwt.Expire
		instance := _jwt.NewJwtInstance(expire, global.Config.Jwt.Secret)

		_, err := instance.ParseToken(split[1])
		if err != nil {
			enum.Result.Error(c, enum.ApiCode.TOKENISVALID, enum.ApiCode.GetMessage(enum.ApiCode.TOKENISVALID))
			c.Abort()
			return
		}
		c.Next()
	}
}
