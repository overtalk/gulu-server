package auth

import (
	"reflect"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type auth struct {
	db module.GameDB
}

func Player(i *module.InternalService) {
	a := auth{db: i.DB}
	i.Gate.POST("/api/v1/login/account", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.Login)
	i.Gate.GET("/api/v1/user/currentUser", a.CurrentPlayer)
	i.Gate.GET("/api/v1/user/notices", a.Notices)
}
