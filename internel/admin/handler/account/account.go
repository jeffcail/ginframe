package account

import (
	"github.com/gin-gonic/gin"
	input2 "github.com/jeffcail/ginframe/internel/admin/input"
	out2 "github.com/jeffcail/ginframe/internel/admin/out"
	"github.com/jeffcail/ginframe/internel/admin/service/account"
	"github.com/jeffcail/ginframe/utils/enum"
	"github.com/spf13/cast"
)

// AccountHandler account handler struct
type AccountHandler struct{}

// List admin account list
func (a *AccountHandler) List(c *gin.Context) {
	param := new(input2.AccountHandlerInput)
	if err := c.Bind(param); err != nil {
		enum.Result.Error(c, enum.ApiCode.PARAMERR, cast.ToString(err))
		return
	}

	count, admins := account.AccountListService(param.PageNum, param.PageSize)
	res := &out2.AccountOut{
		Total: count,
		Data:  admins,
	}
	enum.Result.Success(c, res)
}
