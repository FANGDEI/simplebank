package gateway

type Result struct {
	Succeed bool   `json:"succeed"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

var (
	unknownError = "未知错误, 请联系管理员或稍后重试"
)
