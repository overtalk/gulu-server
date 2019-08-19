package player

import (
	"reflect"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

type player struct {
	db module.GameDB
}

func Player(i *module.InternalService) {
	p := player{db: i.GameDB}
	i.Gate.POST("/v1/user/pvp", reflect.TypeOf(new(protocol.PvpInfo)).Elem(), p.SetPvpInfo)
	i.Gate.POST("/v1/user/info", reflect.TypeOf(new(protocol.BasicInfo)).Elem(), p.SetBasicInfo)
	i.Gate.POST("/v1/user/common", reflect.TypeOf(new(protocol.CommonOP)).Elem(), p.CommonOP)
	i.Gate.POST("/v1/player/query", reflect.TypeOf(new(protocol.PlayerQuery)).Elem(), p.Query)
}
