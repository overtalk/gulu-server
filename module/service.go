package module

import (
	"gitlab.com/SausageShoot/admin-server/utils/gitlab"
)

type InternalService struct {
	DB        GameDB
	Gate      Gate
	GM        GameGM
	UsersConf gitlab.Config
}
