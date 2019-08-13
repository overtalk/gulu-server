package module

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type DB interface {
	GetPlayerByID(id int) (Player, error)
}

type Player interface {
	Apply() error
	SetPvpInfo(info *protocol.PvpInfo) error
	SetBasicInfo(info *protocol.BasicInfo)
}
