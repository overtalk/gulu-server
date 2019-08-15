package gate

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
	. "gitlab.com/SausageShoot/admin-server/utils/trace"
)

var (
	nilEngineErr         = fmt.Errorf("gin engine is nil")
	illegalHttpMethodErr = fmt.Errorf("illegal http method")
)

func (g *gate) AddHandler(httpMethod, relativePath string, ty reflect.Type, handler module.Handler) {
	if g.engine == nil {
		log.Logger.Fatal("add gin route", log.ErrorField(nilEngineErr))
	}

	switch httpMethod {
	case "GET":
		g.engine.GET(relativePath, func(context *gin.Context) {
			context.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: relativePath}))
			body, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				log.Logger.Error("Read GET Request Body", log.ErrorField(err))
				context.String(200, protocol.PostResponse{
					ErrCode: errCode1,
					Msg:     "read get request body error",
				}.Encode())
				return
			}
			context.String(200, handler(context, protocol.GetRequest(ty, body)).Encode())
		})
	case "POST":
		g.engine.POST(relativePath, func(context *gin.Context) {
			context.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: relativePath}))
			body, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				log.Logger.Fatal("Read POST Request Body", log.ErrorField(err))
				context.String(200, protocol.PostResponse{
					ErrCode: errCode2,
					Msg:     "read post request body error",
				}.Encode())
			}
			context.String(200, handler(context, protocol.GetRequest(ty, body)).Encode())
		})
	default:
		log.Logger.Fatal("add gin route",
			log.ErrorField(illegalHttpMethodErr),
			log.Field("method", httpMethod),
			log.Field("handlers", handler))
	}
}
