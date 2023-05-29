package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/core/logger"
	"github.com/jeffcail/ginframe/internel/admin/input"
	out2 "github.com/jeffcail/ginframe/internel/admin/out"
	"github.com/jeffcail/ginframe/utils/enum"
)

// Ping test router
func Ping(c *gin.Context) {
	var param input.PingInput
	err := c.Bind(&param)
	if err != nil {
		enum.Result.Error(c, enum.ApiCode.FAILED, enum.ApiCode.GetMessage(enum.ApiCode.FAILED))
		return
	}

	out := &out2.PingOutput{
		Name:    "test",
		Content: "pong...",
	}

	logger.GinLogger.Info("测试日志能打印出来吗？")

	enum.Result.Success(c, out)
}

// PagePagination pagination
func PagePagination(c *gin.Context) {
	data := []string{"apple", "orange", "banana"}
	list := enum.PageL.Pagination(3, data)
	enum.Result.Success(c, list)
}
