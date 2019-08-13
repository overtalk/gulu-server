package config

import (
	"github.com/spf13/viper"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

const (
	configFile = "config"
	configPath = "$GOPATH/src/gitlab.com/SausageShoot/admin-server/config"
)

func (c *config) getLocalConfig() {
	c.v.SetConfigName(configFile)
	c.v.AddConfigPath(configPath)
	if err := c.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Logger.Fatal("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Logger.Fatal("read config error")
		}
		log.Logger.Fatal("", log.ErrorField(err)) // 读取配置文件失败致命错误
	}
}
