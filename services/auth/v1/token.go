package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

func (a *auth) CurrentPlayer(ctx context.Context, requestMessage interface{}) protocol.Response {
	return protocol.CurrentResponse{
		ErrCode:     1000,
		Name:        "Serati Masss",
		Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Id:          0,
		NotifyCount: 12,
		UnreadCount: 11,
	}
}

func (a *auth) Notices(ctx context.Context, requestMessage interface{}) protocol.Response {
	return protocol.CurrentResponse{
		ErrCode:     1000,
		Name:        "Serati Masss",
		Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Id:          0,
		NotifyCount: 12,
		UnreadCount: 11,
	}
}
