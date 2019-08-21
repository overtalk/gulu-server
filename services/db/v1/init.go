package db

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

type db struct {
	// 所有后台管理用户的账号密码
	users  map[string]module.User
	source module.GMSource
}

func DB(s module.GMSource) *db {
	d := &db{source: s}
	d.loadUsers()

	return d
}

func (a *db) loadUsers() {
	users := new([]module.User)
	data, err := a.source.Fetch("users.json")
	if err != nil {
		log.Logger.Fatal("Read user.json", log.ErrorField(err))
	}
	if err := json.Unmarshal(data, users); err != nil {
		log.Logger.Fatal("Unmarshal user.json", log.ErrorField(err))
	}
	a.users = make(map[string]module.User)
	for _, user := range *users {
		a.users[user.Account] = user
	}
}
