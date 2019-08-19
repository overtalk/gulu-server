package module

// 对游戏数据库的操作

import (
	"gitlab.com/SausageShoot/admin-server/model"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type GameDB interface {
	GetPlayerByID(id int) (GamePlayer, error)
	QueryPlayer(query *protocol.PlayerQuery) (*model.Player, error)
}

type GamePlayer interface {
	Apply() error
	SetPvpInfo(info *protocol.PvpInfo) error
	SetBasicInfo(info *protocol.BasicInfo)
	ClearWeapons(id int) error
	GetAllWeapons(id int) error
	ClearFashions(id int) error
	GetAllFashions(id int) error
	PveUnlock(id int) error
}
