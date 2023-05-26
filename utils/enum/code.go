package enum

type codes struct {
	SUCCESS uint
	FAILED  uint

	MESSAGE map[uint]string
}

var ApiCode = &codes{
	SUCCESS: 2000,
	FAILED:  1000,
}

func InitMapCode() {
	ApiCode.MESSAGE = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED:  "失败",
	}
}

func (c *codes) GetMessage(code uint) string {
	message, ok := c.MESSAGE[code]
	if !ok {
		return "未定义的状态码"
	}
	return message
}
