package enum

type codes struct {
	SUCCESS  uint
	FAILED   uint
	PARAMERR uint

	MESSAGE map[uint]string
}

var ApiCode = &codes{
	SUCCESS:  2000,
	FAILED:   1000,
	PARAMERR: 10001,
}

func InitMapCode() {
	ApiCode.MESSAGE = map[uint]string{
		ApiCode.SUCCESS:  "成功",
		ApiCode.FAILED:   "失败",
		ApiCode.PARAMERR: "参数错误",
	}
}

func (c *codes) GetMessage(code uint) string {
	message, ok := c.MESSAGE[code]
	if !ok {
		return "未定义的状态码"
	}
	return message
}
