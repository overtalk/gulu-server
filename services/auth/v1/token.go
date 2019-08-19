package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

func (a *auth) CurrentPlayer(ctx context.Context) interface{} {
	return protocol.CurrentResponse{
		Response:         protocol.Response{ErrCode: errtable.OkCode},
		Name:             "Serati Masss",
		Avatar:           "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		CurrentAuthority: "admin",
	}
}
