package db

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

type user struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
}

type db struct {
	// 所有后台管理用户的账号密码
	users  map[string]user
	source module.GMSource
}

func DB(s module.GMSource) *db {
	d := &db{source: s}
	d.loadUsers()

	return d
}

func (a *db) loadUsers() {
	users := new([]user)
	data, err := a.source.Fetch("config/user.json")
	if err != nil {
		log.Logger.Fatal("Read user.json", log.ErrorField(err))
	}
	if err := json.Unmarshal(data, users); err != nil {
		log.Logger.Fatal("Unmarshal user.json", log.ErrorField(err))
	}
	a.users = make(map[string]user)
	for _, user := range *users {
		a.users[user.Account] = user
	}
}
