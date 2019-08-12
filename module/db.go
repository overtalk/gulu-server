package module

import (
	"gitlab.com/SausageShoot/admin-server/common"
)

type DB interface {
	GetPlayerByID(id int) (Player, error)
}

type Player interface {
	Apply() error
	SetPvpInfo(info *common.PvpInfo) Player
}
