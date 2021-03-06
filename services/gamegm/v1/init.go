package gamegm

import (
	"sync/atomic"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

type gm struct {
	source      module.GMSource
	weaponList  atomic.Value
	fashionList atomic.Value
	pveNum      int
}

func GM(s module.GMSource) *gm {
	g := &gm{
		source: s,
	}
	// get weapon and fashion
	g.load()
	return g
}

func (g *gm) load() {
	g.loadFashionList()
	g.loadWeaponList()
	g.loadPve()

	log.Logger.Debug("fashion list", log.Field("list", g.GetFashionList()))
	log.Logger.Debug("weapon list", log.Field("list", g.GetWeaponList()))
	log.Logger.Debug("pve num", log.Field("pve num", g.GetPveCount()))
}
