package user

import (
	"github.com/jeffcail/ginframe/server-common/utils/wtime"
	"github.com/jeffcail/ginframe/server-user/daos/account"
	out2 "github.com/jeffcail/ginframe/server-user/out"
	"github.com/jeffcail/ginframe/server-user/ulogger"
	"go.uber.org/zap"
)

func UserListService(pageNum, pageSize int64) (int64, []*out2.Users) {
	count, users, err := account.UserList(pageNum, pageSize)
	if err != nil {
		ulogger.UserLogger.Error("查询管理员账号列表失败, err: ", zap.Error(err))
		return 0, nil
	}

	as := make([]*out2.Users, 0)
	for _, item := range users {
		a := &out2.Users{
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
