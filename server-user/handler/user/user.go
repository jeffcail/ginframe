package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-common/utils/enum"
	input2 "github.com/jeffcail/ginframe/server-user/input"
	out2 "github.com/jeffcail/ginframe/server-user/out"
	"github.com/jeffcail/ginframe/server-user/service/user"
	"github.com/spf13/cast"
)

// UserHandler user handler struct
type UserHandler struct{}

// List user account list
func (a *UserHandler) List(c *gin.Context) {
	param := new(input2.UserHandlerInput)
	if err := c.Bind(param); err != nil {
		enum.Result.Error(c, enum.ApiCode.PARAMERR, cast.ToString(err))
		return
	}

	count, users := user.UserListService(param.PageNum, param.PageSize)
	res := &out2.UserOut{
		Total: count,
		Data:  users,
	}
	enum.Result.Success(c, res)
}
