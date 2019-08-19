package db

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/module"
)

func (a *db) CheckPlayer(username, password string) (int, error) {
	user, isExist := a.users[username]
	if !isExist || user.Password != password {
		return 0, fmt.Errorf("invaild login info")
	}
	return user.ID, nil
}

func (a *db) GetAuth(playerID int) module.Permission {
	for _, v := range a.users {
		if v.ID == playerID {
			return module.Permission(v.Auth)
		}
	}
	return ""
}
