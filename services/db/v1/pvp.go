package db

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/req"
)

// SetPvp set pvp info
func (p *pool) SetPvpInfo(info *req.PvpInfo) error {
	fmt.Println(info)
	return nil
}
