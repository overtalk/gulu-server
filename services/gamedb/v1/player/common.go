package player

import (
	"gitlab.com/SausageShoot/admin-server/utils/parse"
	"gitlab.com/SausageShoot/admin-server/utils/serializer"
)

func (p *player) ClearWeapons(id int) error {
	weaponData, err := encode([]int64{})
	if err != nil {
		return err
	}
	p.pModel.Weapon = weaponData
	p.addUpdate("weapon", nil)
	return nil
}

func (p *player) GetAllWeapons(id int) error {
	weaponData, err := encode(p.gm.GetWeaponList())
	if err != nil {
		return err
	}
	p.pModel.Weapon = weaponData
	p.addUpdate("weapon", weaponData)
	return nil
}

func (p *player) ClearFashions(id int) error {
	fashionData, err := encode([]int64{})
	if err != nil {
		return err
	}
	p.pModel.Fashion = fashionData
	p.addUpdate("fashion", nil)
	return nil
}

func (p *player) GetAllFashions(id int) error {
	fashionData, err := encode(p.gm.GetFashionList())
	if err != nil {
		return err
	}
	p.pModel.Fashion = fashionData
	p.addUpdate("fashion", fashionData)
	return nil
}

func encode(list []int64) (string, error) {
	bytes, err := serializer.Encode(list)
	return string(bytes), err
}

func (p *player) PveUnlock(id int) error {
	type FinishedStage struct {
		ID    string
		Star  int64
		Score int64
	}

	stages := make(map[string]FinishedStage)
	for i := 1; i <= p.gm.GetPveCount(); i++ {
		stageID := "Mission-1-" + parse.String(i)
		stages[stageID] = FinishedStage{
			ID:    stageID,
			Star:  0,
			Score: 0,
		}
	}

	pveData, err := serializer.Encode(stages)
	if err != nil {
		return err
	}

	p.pModel.Pve = string(pveData)
	p.addUpdate("pve", string(pveData))
	return nil
}
