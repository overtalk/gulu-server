package db

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/module"
)

// SetPvp set pvp info
func (p *pool) SetPvpInfo(info *module.PvpInfo) error {
	fmt.Println(info)
	return nil
}
