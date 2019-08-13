package player

import (
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
)

// SetPvp set pvp info
func (p *player) SetPvpInfo(info *protocol.PvpInfo) module.Player {
	// todo: modify pvp data
	return p
}
