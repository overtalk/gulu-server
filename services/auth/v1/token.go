package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/gin-middleware/jwt"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

func (a *auth) CurrentPlayer(ctx context.Context) interface{} {
	resp := protocol.CurrentResponse{
		Response: protocol.Response{ErrCode: errtable.OkCode},
		UserInfo: protocol.UserInfo{
			Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		},
	}
	user, err := a.db.GetUser(parse.Int(ctx.Value(jwt.ID)))
	if err != nil {
		resp.ErrCode = tokenErrCode1
		resp.Msg = "invaild user"
		return resp
	}

	resp.UserInfo.Name = user.Account
	resp.UserInfo.CurrentAuthority = user.Auth
	return resp
}
