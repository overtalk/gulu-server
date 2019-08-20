package main

import (
	"flag"
	"os"
	"os/signal"

	"gitlab.com/SausageShoot/admin-server/services/auth/v1"
	"gitlab.com/SausageShoot/admin-server/services/config/v1"
	"gitlab.com/SausageShoot/admin-server/services/player/v1"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

var (
	port     int
	dev      bool
	logLevel string
	distPath string
)

func main() {
	flag.IntVar(&port, "port", 9000, "server port")
	flag.BoolVar(&dev, "dev", false, "`dev` mode use local file config")
	flag.StringVar(&logLevel, "log-level", "info", "log level, optional( debug | info | warn | error | dpanic | panic | fatal), default is info")
	flag.StringVar(&distPath, "dist-path", "/", "static dist path")
	flag.Parse()

	// init logger
	log.InitLogger(logLevel)

	internalService := config.Config(dev).InternalService(port)
	//internalService.Gate.AddStatic("/", distPath, true)

	auth.Auth(internalService)
	player.Player(internalService)

	// start the server
	internalService.Gate.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	internalService.Gate.Stop()
}
