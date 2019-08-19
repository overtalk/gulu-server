package gamegm

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

// get pve list from gm source
func (g *gm) loadPve() {
	fileName := "stage.txt"
	data, err := g.source.Fetch(fileName)
	if err != nil {
		log.Logger.Fatal("Fetch Pve Stage List",
			log.ErrorField(err))
	}

	stages := make(map[string]interface{})
	if err := json.Unmarshal(data, &stages); err != nil {
		log.Logger.Fatal("Unmarshal Stage List",
			log.ErrorField(err))
	}

	g.pveNum = len(stages)
}
func (g *gm) GetPveCount() int {
	return g.pveNum
}
