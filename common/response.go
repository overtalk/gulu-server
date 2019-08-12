package common

import (
	"encoding/json"
)

type Response interface {
	Encode() string
}

// post response from client
type Post struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

func (p Post) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		p.ErrCode = 1001
		p.Msg = "反序列化错误"
	}
	return string(d)
}
