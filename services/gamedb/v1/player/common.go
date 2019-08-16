package player

import (
	"gitlab.com/SausageShoot/admin-server/utils/serializer"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

// 1: 获取所有武器
// 2: 删除所有武器
// 3: 获取所有时装
// 4: 清除所有时装
func (p *player) CommonOP(info *protocol.CommonOP) error {
	switch info.OP {
	case 1:
		weaponData, err := encode(p.gm.GetWeaponList())
		if err != nil {
			return err
		}
		p.pModel.Weapon = weaponData
		p.addUpdate("weapon", weaponData)
	case 2:
		weaponData, err := encode([]int64{})
		if err != nil {
			return err
		}
		p.pModel.Weapon = weaponData
		p.addUpdate("weapon", weaponData)
	case 3:
		fashionData, err := encode(p.gm.GetFashionList())
		if err != nil {
			return err
		}
		p.pModel.Fashion = fashionData
		p.addUpdate("fashion", fashionData)
	case 4:
		fashionData, err := encode([]int64{})
		if err != nil {
			return err
		}
		p.pModel.Fashion = fashionData
		p.addUpdate("fashion", fashionData)
	}
	return nil
}

func encode(list []int64) (string, error) {
	bytes, err := serializer.Encode(list)
	return string(bytes), err
}
