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
	ID      int    `json:"id"`
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

// login response from client
type LoginResponse struct {
	Status           string `json:"status"`
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
	Token            string `json:"token"`
}

func (p LoginResponse) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		log.Logger.Error("LoginResponse Encode",
			log.ErrorField(err),
			log.Field("resp", p))
		return responseEncodeErr
	}
	return string(d)
}

type CurrentResponse struct {
	ErrCode     int    `json:"errcode"`
	Msg         string `json:"msg"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Id          int    `json:"id"`
	NotifyCount int    `json:"notifyCount"`
	UnreadCount int    `json:"unreadCount"`
}

func (p CurrentResponse) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		log.Logger.Error("CurrentResponse Encode",
			log.ErrorField(err),
			log.Field("resp", p))
		return responseEncodeErr
	}
	return string(d)
}
