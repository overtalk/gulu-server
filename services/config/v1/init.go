package config

import (
	"github.com/spf13/viper"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/services/db/v1"
	"gitlab.com/SausageShoot/admin-server/services/gate/v1"
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
	return &module.InternalService{
		DB: db.Pool(db.Config{
			Username: c.v.GetString(mySQLUsername),
			Password: c.v.GetString(mySQLPassword),
			Host:     c.v.GetString(mySQLHost),
			DBName:   c.v.GetString(mySQLDBName),
			Port:     c.v.GetInt(mySQLPort),
		}),
		Gate: gate.Gate(port),
	}
}
