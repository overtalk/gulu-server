package hub000

import (
	"gitlab.com/SausageShoot/admin-server/utils/parse"
	"gitlab.com/SausageShoot/admin-server/utils/serializer"
)

type Arena struct {
	Cup     map[int64]int64 // old cup
	CupHide map[int64]int64
	Lock    int64
	MaxCup  int64
	Version int64
}

func DecodeArena(data string) Arena {
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

func EncodeArena(arena Arena) (string, error) {
	newArena, err := serializer.Encode(arena)
	if err != nil {
		return "", err
	}
	return string(newArena), nil
}
