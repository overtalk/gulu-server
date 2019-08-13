package protocol

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

const responseEncodeErr = `{"errcode":1003,"msg":"encode response error"}`

type Response interface {
	Encode() string
}

// post response from client
type PostResponse struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

func (p PostResponse) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		log.Logger.Error("PsostResponse Encode",
			log.ErrorField(err),
			log.Field("resp", p))
		return responseEncodeErr
	}
	return string(d)
}
