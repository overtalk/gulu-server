package auth

import (
	"context"
	"fmt"

	//"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (a *auth) Login(ctx context.Context, requestMessage interface{}) protocol.Response {
	req, ok := requestMessage.(*protocol.LoginReq)
	//resp := protocol.PostResponse{ErrCode: errtable.OkCode}

	if !ok {
		log.Logger.Error("Convert", log.Field("request", requestMessage))
		return protocol.PostResponse{
			ErrCode: authErrCode1,
			Msg:     "convert CommonOP request error",
		}
	}

	fmt.Println(req)

	return protocol.LoginResponse{
		Status:           "ok",
		CurrentAuthority: "admin",
		Token:            "token",
	}
}
