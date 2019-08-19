package db

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/module"
)

func (a *db) CheckUser(username, password string) (int64, error) {
	user, isExist := a.users[username]
	if !isExist || user.Password != password {
		return 0, fmt.Errorf("invaild login info")
	}
	return user.ID, nil
}

func (a *db) GetUser(id int64) (*module.User, error) {
	for _, v := range a.users {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("invaild player id")
}
