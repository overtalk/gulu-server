package main

import (
	"flag"
	"os"
	"os/signal"

	"gitlab.com/SausageShoot/admin-server/services/gate/v1"
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

var (
	port     int
	logLevel string
)

func main() {
	flag.IntVar(&port, "port", 9000, "server port")
	flag.StringVar(&logLevel, "log-level", "info", "log level, optional( debug | info | warn | error | dpanic | panic | fatal), default is info")
	flag.Parse()

	// init logger
	logger.InitLogger(logLevel)

	g := gate.Gate(port)

	// start the server
	g.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	g.Stop()
}
