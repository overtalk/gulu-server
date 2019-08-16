package openid_test

import (
	"testing"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/services/openid/v1"
	"gitlab.com/SausageShoot/admin-server/utils/mysql"
)

func getConn(t *testing.T) module.GameOpenIdDB {
	c := mysql.Config{
		Username: "sausage",
		Password: "sausage_shooter",
		Host:     "172.26.32.12",
		DBName:   "sausage_shooter",
		Port:     3306,
		MaxConn:  0,
	}
	return openid.NewOpenIdDB(c)
}

func TestConnect(t *testing.T) {
	getConn(t)
}
