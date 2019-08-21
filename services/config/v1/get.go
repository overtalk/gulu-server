package config

import (
	"bytes"
	"os"

	"github.com/spf13/viper"

	"gitlab.com/SausageShoot/admin-server/utils/gitlab"
	"gitlab.com/SausageShoot/admin-server/utils/log"
	"gitlab.com/SausageShoot/admin-server/utils/parse"
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

func (c *config) getRemoteConfig() {
	keys := []string{
		"GITLAB_TOKEN",
		"GITLAB_REF",
		"GITLAB_PID",
		"GITLAB_URL",
		"GITLAB_PATH",
	}
	envs := make(map[string]string)
	for _, key := range keys {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			log.Logger.Fatal("Missing Env", log.Field("env-key", key))
		}
		envs[key] = value
	}

	configData, err := gitlab.Catcher(gitlab.Config{
		Token: envs["GITLAB_TOKEN"],
		Ref:   envs["GITLAB_REF"],
		Pid:   int(parse.Int(envs["GITLAB_PID"])),
		Url:   envs["GITLAB_URL"],
	}).Fetch(envs["GITLAB_PATH"])
	if err != nil {
		log.Logger.Fatal("Get File Data", log.ErrorField(err))
	}

	c.v.SetConfigType("yaml")
	if err := c.v.ReadConfig(bytes.NewBuffer(configData)); err != nil {
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
