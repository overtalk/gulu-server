package db

import (
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/services/openid/v1"
	"gitlab.com/SausageShoot/admin-server/utils/mysql"
)

var (
	playerTable = "player"
)

type pool struct {
	openidDB module.GameOpenIdDB
	gm       module.GameGM
}

// Pool is the constructor
func Pool(c mysql.Config, gm module.GameGM) *pool {
	return &pool{
		openidDB: openid.NewOpenIdDB(c),
		gm:       gm,
	}
}
