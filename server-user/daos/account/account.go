package account

import (
	"github.com/jeffcail/ginframe/server-user/models"
)

func UserList(pageNum, pageSize int64) (int64, []*models.User, error) {
	a := new(models.User)
	return a.List(pageNum, pageSize)
}
