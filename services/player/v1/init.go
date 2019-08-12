package player

import (
	"gitlab.com/SausageShoot/admin-server/module"
)

type player struct {
	db module.DB
}

func Player(i *module.InternalService) {
	p := player{
		db: i.DB,
	}

	i.Gate.Add("POST", "/v1/user/pvp", p.SetPvpInfo)
	i.Gate.Add("POST", "/v1/user/info", p.SetBasicInfo)
	i.Gate.Add("POST", "/v1/user/common", p.CommonOp)
}
