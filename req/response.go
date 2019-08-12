package req

import (
	"encoding/json"
)

type PostResponse struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

func (p PostResponse) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		p.ErrCode = 1001
		p.Msg = "反序列化错误"
	}
	return string(d)
}
