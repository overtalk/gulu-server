package protocol

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

// LoginReq is login request struct
type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// login response from client
type LoginResponse struct {
	Status string `json:"status"`
	Type   string `json:"type"`
	Token  string `json:"token"`
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
	ErrCode          int    `json:"errcode"`
	Msg              string `json:"msg"`
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	Id               int    `json:"id"`
	CurrentAuthority string `json:"currentAuthority"`
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
