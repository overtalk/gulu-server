package module

// 对游戏数据库的操作

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type GameDB interface {
	GetPlayerByID(id int) (GamePlayer, error)
	QueryPlayer(query *protocol.PlayerQuery) error
}

type GamePlayer interface {
	Apply() error
	SetPvpInfo(info *protocol.PvpInfo) error
	CommonOP(info *protocol.CommonOP) error
	SetBasicInfo(info *protocol.BasicInfo)
}