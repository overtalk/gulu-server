package auth

import (
	"encoding/json"
	"reflect"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/gitlab"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

type user struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
	Name     string `json:"name"`
}

type auth struct {
	users  []user
	config gitlab.Config
	db     module.GameDB
}

func Auth(i *module.InternalService) {
	a := auth{db: i.DB, config: i.UsersConf}
	a.loadUsers()

	i.Gate.POST("/api/v1/login/account", reflect.TypeOf(new(protocol.LoginReq)).Elem(), a.Login)
	i.Gate.GET("/api/v1/user/currentUser", a.CurrentPlayer)
	i.Gate.GET("/api/v1/user/notices", a.Notices)
}

func (a *auth) loadUsers() {
	users := new([]user)
	data, err := gitlab.Catcher(a.config).Fetch("config/user.json")
	if err != nil {
		log.Logger.Fatal("Read user.json", log.ErrorField(err))
	}
	if err := json.Unmarshal(data, users); err != nil {
		log.Logger.Fatal("Unmarshal user.json", log.ErrorField(err))
	}
	a.users = *users
}
