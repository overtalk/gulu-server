package player

import (
	"gitlab.com/SausageShoot/admin-server/errtable"
)

const (
	m = errtable.Player
)

var (
	// common op
	commonErrCode1 = errtable.GenCode(errtable.System, m, 1)
	commonErrCode2 = errtable.GenCode(errtable.System, m, 2)
	commonErrCode3 = errtable.GenCode(errtable.System, m, 3)
	commonErrCode4 = errtable.GenCode(errtable.System, m, 4)
	// info
	infoErrCode1 = errtable.GenCode(errtable.System, m, 5)
	infoErrCode2 = errtable.GenCode(errtable.System, m, 6)
	infoErrCode3 = errtable.GenCode(errtable.System, m, 7)
	// pvp op
	pvpErrCode1 = errtable.GenCode(errtable.System, m, 8)
	pvpErrCode2 = errtable.GenCode(errtable.System, m, 9)
	pvpErrCode3 = errtable.GenCode(errtable.System, m, 10)
	pvpErrCode4 = errtable.GenCode(errtable.System, m, 11)
	// query
	queryErrCode1 = errtable.GenCode(errtable.System, m, 12)
	queryErrCode2 = errtable.GenCode(errtable.System, m, 13)
)
