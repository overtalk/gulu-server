package gate

import (
	"fmt"

	"gitlab.com/SausageShoot/admin-server/errtable"
)

// common error
var nilEngineErr = fmt.Errorf("gin engine is nil")

const (
	m = errtable.Gate
)

var (
	errCode1 = errtable.GenCode(errtable.System, m, 1)
	errCode2 = errtable.GenCode(errtable.System, m, 2)
)
