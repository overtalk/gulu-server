package gate

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
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

func (g *gate) Add1(httpMethod, relativePath string, handlers module.Handler) {
	if g.engine == nil {
		logger.Logger.Fatal("add gin route", logger.ErrorField(nilEngineErr))
	}

	switch httpMethod {
	case "GET":
		g.engine.GET(relativePath, func(context *gin.Context) {
			context.String(200, handlers(1).Encode())
		})
	case "POST":
		g.engine.POST(relativePath, func(context *gin.Context) {
			context.String(200, handlers(1).Encode())
		})
	default:
		logger.Logger.Fatal("add gin route",
			logger.ErrorField(illegalHttpMethodErr),
			logger.Field("method", httpMethod),
			logger.Field("handlers", handlers))
	}

}
