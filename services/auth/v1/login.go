package auth

import (
	"context"
	//"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (a *auth) Login(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.LoginReq)
	resp := protocol.LoginResponse{Status: "ok", Type: "account"}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: authErrCode1,
			Msg:     "convert CommonOP request error",
		}
	}

	user, isExist := a.users[req.Username]
	if !isExist || user.Password != req.Password {
		resp.Status = "error"
	} else {
		resp.CurrentAuthority = user.Auth
		resp.Token = user.Account + "-token"
	}

	return resp
}
