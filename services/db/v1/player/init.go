package player

import (
	"gitlab.com/SausageShoot/admin-server/model"
)

type player struct {
}

func Player(pl model.Player) *player {
	return &player{}
}

// write data to db
func (p *player) Apply() error {
	// todo: apply data to db
	return nil
}
