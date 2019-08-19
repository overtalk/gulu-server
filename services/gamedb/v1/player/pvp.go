package player

import (
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/hub000"
)

// SetPvp set pvp info
func (p *player) SetPvpInfo(info *protocol.PvpInfo) error {
	// decode
	arena := hub000.DecodeArena(p.pModel.Arena)

	arena.Version = 1
	// set lock
	arena.Lock = int64(info.Arena)
	for k, v := range info.Socre {
		arena.CupHide[int64(k)+1] = int64(v)
	}

	// encode arena
	data, err := hub000.EncodeArena(arena)
	if err != nil {
		return err
	}
	p.pModel.Arena = data

	// add update
	p.addUpdate("arena", data)
	return nil
}
