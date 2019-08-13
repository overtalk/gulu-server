package module

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type DB interface {
	GetPlayerByID(id int) (Player, error)
	QueryPlayer(query *protocol.PlayerQuery) error
}

type Player interface {
	Apply() error
	SetPvpInfo(info *protocol.PvpInfo) error
	CommonOP(info *protocol.CommonOP) error
	SetBasicInfo(info *protocol.BasicInfo)
}
