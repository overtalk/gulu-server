package auth

import (
	"gitlab.com/SausageShoot/admin-server/errtable"
)

const (
	m = errtable.Auth
)

var (
	// auth
	authErrCode1 = errtable.GenCode(errtable.System, m, 1)
	authErrCode2 = errtable.GenCode(errtable.System, m, 2)

	// token
	tokenErrCode1 = errtable.GenCode(errtable.System, m, 3)
)
