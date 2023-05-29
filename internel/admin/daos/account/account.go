package account

import (
	"github.com/jeffcail/ginframe/internel/admin/models"
)

func AccountList(pageNum, pageSize int64) (int64, []*models.Admin, error) {
	a := new(models.Admin)
	return a.List(pageNum, pageSize)
}
