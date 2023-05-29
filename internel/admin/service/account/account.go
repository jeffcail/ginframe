package account

import (
	"github.com/jeffcail/ginframe/core/logger"
	"github.com/jeffcail/ginframe/internel/admin/daos/account"
	out2 "github.com/jeffcail/ginframe/internel/admin/out"
	"github.com/jeffcail/ginframe/utils/wtime"
	"go.uber.org/zap"
)

func AccountListService(pageNum, pageSize int64) (int64, []*out2.Accounts) {
	count, admins, err := account.AccountList(pageNum, pageSize)
	if err != nil {
		logger.GinLogger.Error("查询管理员账号列表失败, err: ", zap.Error(err))
		return 0, nil
	}

	as := make([]*out2.Accounts, 0)
	for _, item := range admins {
		a := &out2.Accounts{
			Id:        item.Id,
			Username:  item.Username,
			Nickname:  item.Nickname,
			Phone:     item.Phone,
			Email:     item.Email,
			Gender:    item.Gender,
			RoleId:    item.RoleId,
			Enable:    item.Enable,
			CreatedAt: wtime.WTtime.FormatTime(item.CreatedAt),
			UpdatedAt: wtime.WTtime.FormatTime(item.UpdatedAt),
		}
		as = append(as, a)
	}

	return count, as
}
