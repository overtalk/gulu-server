package player

import (
	"gitlab.com/SausageShoot/admin-server/utils/serializer"

	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
)

type Arena struct {
	Cup     map[int64]int64 // old cup
	CupHide map[int64]int64
	Lock    int64
	MaxCup  int64
	Version int64
}

// SetPvp set pvp info
func (p *player) SetPvpInfo(info *protocol.PvpInfo) error {
	// judge the input is legal
	if !(info.Arena > 0 && info.Arena < 10) && !(info.Cup > 0) {
		return nil
	}

	// decode
	arena := decodeArena(p.pModel.Arena)

	// set lock
	if info.Arena > 0 && info.Arena < 10 {
		arena.Lock = int64(info.Arena)
	}
	// set cup
	if info.Cup > 0 {
		arena.MaxCup = int64(info.Cup)
	}

	// encode arena
	data, err := encodeArena(arena)
	if err != nil {
		return err
	}
	p.pModel.Arena = data

	// add update
	p.addUpdate("arena", data)
	return nil
}

func decodeArena(data string) Arena {
	var arena Arena
	if err := serializer.Decode(parse.Bytes(data), &arena); err != nil {
		arena = Arena{Cup: make(map[int64]int64), CupHide: make(map[int64]int64), Lock: 0, Version: 0}
	}
	if arena.Cup == nil {
		arena.Cup = make(map[int64]int64)
	}
	if arena.CupHide == nil {
		arena.CupHide = make(map[int64]int64)
	}
	switch arena.Version {
	case 0:
		for i := int64(1); i <= 10; i++ {
			arena.CupHide[i] = 1000
		}
	}
	arena = Arena{Cup: arena.CupHide, CupHide: arena.CupHide, Lock: arena.Lock, Version: arena.Version}
	return arena
}

func encodeArena(arena Arena) (string, error) {
	newArena, err := serializer.Encode(arena)
	if err != nil {
		return "", err
	}
	return string(newArena), nil
}
