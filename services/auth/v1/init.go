package auth

import (
	"reflect"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type auth struct {
	db     module.DB
	gameDB module.GameDB
}

func Auth(i *module.InternalService) {
	a := auth{gameDB: i.GameDB, db: i.DB}

	i.Gate.POST("/v1/login/account", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.Login)
	i.Gate.GET("/v1/user/currentUser", a.CurrentPlayer)
}
