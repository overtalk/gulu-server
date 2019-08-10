package gate

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

var (
	nilEngineErr         = fmt.Errorf("gin engine is nil")
	illegalHttpMethodErr = fmt.Errorf("illegal http method")
)

func (g *gate) Add(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	if g.engine == nil {
		logger.Logger.Fatal("add gin route", logger.ErrorField(nilEngineErr))
	}

	switch httpMethod {
	case "GET":
		g.engine.GET(relativePath, handlers...)
	case "POST":
		g.engine.POST(relativePath, handlers...)
	default:
		logger.Logger.Fatal("add gin route",
			logger.ErrorField(illegalHttpMethodErr),
			logger.Field("method", httpMethod),
			logger.Field("handlers", handlers))
	}

}
