package gm

import (
	"encoding/json"

	"gitlab.com/SausageShoot/admin-server/utils/log"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

// get fashion list from gm source
func (g *gm) loadFashionList() {
	fileName := "FashionList.txt"
	data, err := g.source.Fetch(fileName)
	if err != nil {
		log.Logger.Fatal("Fetch Fashion List",
			log.ErrorField(err))
	}

	fashions := make(map[string]interface{})
	if err := json.Unmarshal(data, &fashions); err != nil {
		log.Logger.Fatal("Unmarshal Fashion List",
			log.ErrorField(err))
	}

	var fashionList []int64
	for k, _ := range fashions {
		fashionList = append(fashionList, parse.Int(k))
	}

	g.fashionList.Store(fashionList)
}

func (g *gm) GetFashionList() []int64 {
	return g.fashionList.Load().([]int64)
}
