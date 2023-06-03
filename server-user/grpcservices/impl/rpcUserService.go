package impl

import (
	"github.com/jeffcail/ginframe/server-user/models"
	"strconv"
)

// GetUserInfo 获取管理员账号信息
func GetUserInfo(id string) *models.User {
	aid, _ := strconv.Atoi(id)

	m := new(models.User)
	info, err := m.FindUserById(int64(aid))
	if err != nil {
		return nil
	}
	return info
}
