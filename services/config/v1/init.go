package config

import (
	"github.com/spf13/viper"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/services/gamedb/v1"
	"gitlab.com/SausageShoot/admin-server/services/gamegm/v1"
	"gitlab.com/SausageShoot/admin-server/services/gate/v1"
	"gitlab.com/SausageShoot/admin-server/utils/gitlab"
	"gitlab.com/SausageShoot/admin-server/utils/mysql"
)

type config struct {
	v *viper.Viper
}

func Config() *config {
	c := &config{v: viper.New()}

	// get config
	c.getLocalConfig()

	return c
}

func (c *config) InternalService(port int) *module.InternalService {
	gm := gamegm.GM(gitlab.Catcher(gitlab.Config{
		Token: c.v.GetString(gmToken),
		Ref:   c.v.GetString(gmBranch),
		Pid:   c.v.GetInt(gmPid),
		Url:   c.v.GetString(gmUrl),
	}))
	return &module.InternalService{
		UsersConf: gitlab.Config{
			Token: c.v.GetString(userToken),
			Ref:   c.v.GetString(userBranch),
			Pid:   c.v.GetInt(userPid),
			Url:   c.v.GetString(userUrl),
		},
		DB: gamedb.Pool(mysql.Config{
			Username: c.v.GetString(mySQLUsername),
			Password: c.v.GetString(mySQLPassword),
			Host:     c.v.GetString(mySQLHost),
			DBName:   c.v.GetString(mySQLDBName),
			Port:     c.v.GetInt(mySQLPort),
		}, gm),
		Gate: gate.Gate(port),
		GM:   gm,
	}
}
