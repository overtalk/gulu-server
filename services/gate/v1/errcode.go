package gate

import (
	"gitlab.com/SausageShoot/admin-server/errtable"
)

const (
	m = errtable.Gate
)

var (
	errCode1 = errtable.GenCode(errtable.System, m, 1)
	errCode2 = errtable.GenCode(errtable.System, m, 2)
)
