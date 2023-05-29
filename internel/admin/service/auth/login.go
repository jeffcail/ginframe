package auth

import (
	"errors"
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/core/logger"
	"github.com/jeffcail/ginframe/internel/admin/daos/auth"
	"github.com/jeffcail/ginframe/internel/admin/input"
	out2 "github.com/jeffcail/ginframe/internel/admin/out"
	_jwt "github.com/jeffcail/ginframe/pkg/jwt"
	"github.com/jeffcail/ginframe/utils/encry"
)

func LoginService(param *input.LoginInput) (*out2.LoginOut, error) {
	admin, err := auth.FindAdminByUsername(param.Username)
	if err != nil {
		logger.GinLogger.Error(fmt.Sprintf("username: %s 账号不存在", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 账号不存在", param.Username))
	}

	err = encry.ComparePassword(admin.Password, param.Password)
	if err != nil {
		logger.GinLogger.Error(fmt.Sprintf("username: %s 密码错误", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 密码错误", param.Username))
	}

	claims := _jwt.NewJwtInstance(global.Config.Jwt.Expire, global.Config.Jwt.Secret)
	token, err := claims.GenerateToken()
	if err != nil {
		logger.GinLogger.Error(fmt.Sprintf("username: %s 签发token失败", param.Username))
		return nil, errors.New(fmt.Sprintf("username: %s 签发token失败", param.Username))
	}

	out := new(out2.LoginOut)
	out.Token = token
	out.Username = admin.Username

	return out, nil

}
