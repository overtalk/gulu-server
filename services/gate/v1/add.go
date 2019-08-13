package gate

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/errtable"
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

var (
	nilEngineErr         = fmt.Errorf("gin engine is nil")
	illegalHttpMethodErr = fmt.Errorf("illegal http method")
)

func (g *gate) AddHandler(httpMethod, relativePath string, ty reflect.Type, handler module.Handler) {
	if g.engine == nil {
		logger.Logger.Fatal("add gin route", logger.ErrorField(nilEngineErr))
	}

	switch httpMethod {
	case "GET":
		g.engine.GET(relativePath, func(context *gin.Context) {
			body, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				logger.Logger.Fatal("Read GET Request Body", logger.ErrorField(err))
				context.String(200, errtable.ReadBodyErr.Encode())
			}
			context.String(200, handler(protocol.GetRequest(ty, body)).Encode())
		})
	case "POST":
		g.engine.POST(relativePath, func(context *gin.Context) {
			body, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				logger.Logger.Fatal("Read POST Request Body", logger.ErrorField(err))
				context.String(200, errtable.ReadBodyErr.Encode())
			}
			context.String(200, handler(protocol.GetRequest(ty, body)).Encode())
		})
	default:
		logger.Logger.Fatal("add gin route",
			logger.ErrorField(illegalHttpMethodErr),
			logger.Field("method", httpMethod),
			logger.Field("handlers", handler))
	}
}
