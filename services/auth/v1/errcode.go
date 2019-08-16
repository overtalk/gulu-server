package auth

import (
	"gitlab.com/SausageShoot/admin-server/errtable"
)

const (
	m = errtable.Auth
)

var (
	// common op
	authErrCode1 = errtable.GenCode(errtable.System, m, 1)
	authErrCode2 = errtable.GenCode(errtable.System, m, 2)
	authErrCode3 = errtable.GenCode(errtable.System, m, 3)
	authErrCode4 = errtable.GenCode(errtable.System, m, 4)
)
