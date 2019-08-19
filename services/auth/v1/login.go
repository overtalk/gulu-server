package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/gin-middleware/jwt"
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

	playerID, err := a.db.CheckUser(req.Username, req.Password)
	if err != nil {
		resp.Status = "error"
	} else {
		token, err := jwt.NewJWT().CreateToken(jwt.CustomClaims{ID: int(playerID)})
		if err != nil {
			resp.ErrCode = authErrCode2
			resp.Msg = "gen jwt error"
			return resp
		}
		resp.Token = token
	}

	return resp
}
