package enum

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var (
	Result *result
	once   sync.Once
	PageL  *pageList
)

type result struct {
	Time time.Time   `json:"time"`
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功返回
func (res *result) Success(c *gin.Context, data ...interface{}) {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}
	res.Time = time.Now()
	res.Code = ApiCode.SUCCESS
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = d
	c.JSON(http.StatusOK, res)
}

// Error 失败返回
func (res *result) Error(c *gin.Context, code uint, msg string) {
	res.Time = time.Now()
	res.Code = code
	res.Msg = ApiCode.GetMessage(code)
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}

type pageList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// Pagination 分页
func (p *pageList) Pagination(count int64, list interface{}) *pageList {
	return &pageList{
		Total: count,
		List:  list,
	}
}

func init() {
	InitMapCode()
	once.Do(func() {
		Result = &result{}
		PageL = &pageList{}
	})
}
