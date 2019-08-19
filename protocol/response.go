package protocol

// Response : 通用的返回，包含errcode和msg
type Response struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
}
