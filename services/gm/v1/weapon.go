package gm

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/utils/log"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

// get weapon list from gm source
func (g *gm) loadWeaponList() {
	fileName := "WeaponList.txt"
	data, err := g.source.Fetch(fileName)
	if err != nil {
		log.Logger.Fatal("Fetch weapon List",
			log.ErrorField(err))
	}

	weapons := make(map[string]interface{})
	if err := json.Unmarshal(data, &weapons); err != nil {
		log.Logger.Fatal("Unmarshal weapon List",
			log.ErrorField(err))
	}

	var weaponList []int64
	for k, _ := range weapons {
		weaponList = append(weaponList, parse.Int(k))
	}

	g.weaponList.Store(weaponList)
}

func (g *gm) GetWeaponList() []int64 {
	return g.weaponList.Load().([]int64)
}
