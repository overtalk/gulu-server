package gate

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

type Gate struct {
	port   int
	engine gin.Engine
}

func (g *Gate) Start() {
	go func(g *Gate) {
		if err := g.engine.Run(fmt.Sprintf(":%d", g.port)); err != nil {
			logger.Logger.Error("start server", logger.ErrorField(err))
		}
	}(g)
}

func (g *Gate) Stop() {

}
