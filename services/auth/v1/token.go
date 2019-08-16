package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

func (a *auth) CurrentPlayer(ctx context.Context) protocol.Response {
	return protocol.CurrentResponse{
		ErrCode:          errtable.OkCode,
		Name:             "Serati Masss",
		Avatar:           "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Id:               0,
		CurrentAuthority: "admin",
	}

	// get auth by token

	//return resp

}
