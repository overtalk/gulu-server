package main

import (
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func main() {
	logger.InitLogger("info")
	logger.Logger.Info("ok")
}
