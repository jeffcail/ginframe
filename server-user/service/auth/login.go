package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_jwt "github.com/jeffcail/ginframe/server-common/pkg/jwt"
	"github.com/jeffcail/ginframe/server-common/utils/encry"
	"github.com/jeffcail/ginframe/server-user/daos/auth"
	"github.com/jeffcail/ginframe/server-user/global"
	"github.com/jeffcail/ginframe/server-user/input"
	out2 "github.com/jeffcail/ginframe/server-user/out"
	"github.com/jeffcail/ginframe/server-user/ulogger"
	"time"
)

func LoginService(param *input.LoginInput) (*out2.LoginOut, error) {
	u, err := auth.FindUserByUsername(param.Username)
	if err != nil {
		ulogger.UserLogger.Error(fmt.Sprintf("username: %s 账号不存在", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 账号不存在", param.Username))
	}

	err = encry.ComparePassword(u.Password, param.Password)
	if err != nil {
		ulogger.UserLogger.Error(fmt.Sprintf("username: %s 密码错误", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 密码错误", param.Username))
	}

	expire := time.Hour * global.Config.Jwt.Expire
	instance := _jwt.NewJwtInstance(expire, global.Config.Jwt.Secret)
	claims := &_jwt.JwtClaims{
		ID:             u.Id,
		Username:       u.Username,
		StandardClaims: jwt.StandardClaims{},
	}

	token, err := instance.GenerateToken(claims)
	if err != nil {
		ulogger.UserLogger.Error(fmt.Sprintf("username: %s 签发token失败", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 签发token失败", param.Username))
	}

	out := new(out2.LoginOut)
	out.Token = token
	out.Username = u.Username

	return out, nil

}
