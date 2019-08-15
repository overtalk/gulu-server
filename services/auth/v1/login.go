package auth

import (
	"context"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

func (a *auth) Login(ctx context.Context, requestMessage interface{}) protocol.Response {
	return protocol.LoginResponse{
		Status: "ok",
		//Type:             "accoun",
		CurrentAuthority: "admin",
		Token:            "token",
	}
}
