package auth

import (
	"context"
	"fmt"

	"gitlab.com/SausageShoot/admin-server/errtable"
	//"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (a *auth) Login(ctx context.Context, requestMessage interface{}) interface{} {
	req, ok := requestMessage.(*protocol.LoginReq)
	resp := protocol.LoginResponse{Response: protocol.Response{ErrCode: errtable.OkCode}, Status: "ok", Type: "account"}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		resp.ErrCode = authErrCode1
		resp.Msg = "convert Login request error"
		return resp
	}

	playerID, err := a.db.CheckPlayer(req.Username, req.Password)
	if err != nil {
		resp.Status = "error"
	} else {
		fmt.Println(playerID)
		// todo: gen token
		resp.Token = req.Username + "-token"
	}

	return resp
}
