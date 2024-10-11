// 状态码
// author daniel

package result

// Codes 定义状态
type Codes struct {
	SUCCESS  uint
	FAILED   uint
	Message  map[uint]string
	REQUIRED uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:  200,
	FAILED:   501,
	REQUIRED: 400,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:  "成功",
		ApiCode.FAILED:   "失败",
		ApiCode.REQUIRED: "缺少必要参数",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
