package player

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
)

func (p *player) SetBasicInfo(info *protocol.BasicInfo) {
	if info.Diamond > 0 {
		p.pModel.Diamond = info.Diamond
		p.addUpdate("diamond", info.Diamond)
	}

	if info.Gold > 0 {
		p.pModel.Gold = info.Gold
		p.addUpdate("gold", info.Gold)
	}

	if info.Experience > 0 {
		p.pModel.Experience = info.Experience
		p.addUpdate("experience", info.Experience)
	}

	if info.Strength > 0 {
		p.pModel.Strength = info.Strength
		p.addUpdate("strength", info.Strength)
	}
}
