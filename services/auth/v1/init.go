package auth

import (
	"reflect"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type auth struct {
	db module.DB
}

func Player(i *module.InternalService) {
	a := auth{db: i.DB}
	i.Gate.AddHandler("POST", "/api/v1/login/account", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.Login)
	i.Gate.AddHandler("GET", "/api/v1/user/currentUser", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.CurrentPlayer)
	i.Gate.AddHandler("GET", "/api/v1/user/notices", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.Notices)
}
