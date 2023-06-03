package enum

type codes struct {
	SUCCESS          uint
	FAILED           uint
	PARAMERR         uint
	TOKENISNOTEXISTS uint
	TOKENISVALID     uint
	USERISDISABLE    uint

	MESSAGE map[uint]string
}

var ApiCode = &codes{
	SUCCESS:          2000,
	FAILED:           1000,
	PARAMERR:         10001,
	TOKENISNOTEXISTS: 401,
	TOKENISVALID:     402,
	USERISDISABLE:    403,
}

func InitMapCode() {
	ApiCode.MESSAGE = map[uint]string{
		ApiCode.SUCCESS:          "成功",
		ApiCode.FAILED:           "失败",
		ApiCode.PARAMERR:         "参数错误",
		ApiCode.TOKENISNOTEXISTS: "token不存在,非法请求",
		ApiCode.TOKENISVALID:     "token无效,非法请求",
		ApiCode.USERISDISABLE:    "账号被禁用，请联系管理员处理",
	}
}

func (c *codes) GetMessage(code uint) string {
	message, ok := c.MESSAGE[code]
	if !ok {
		return "未定义的状态码"
	}
	return message
}
