package gate

import (
	"reflect"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (g *gate) POST(relativePath string, ty reflect.Type, handler module.POSTHandler) {
	if g.engine == nil {
		log.Logger.Fatal("add gin post route",
			log.Field("path", relativePath),
			log.Field("request", ty),
			log.Field("handler", handler),
			log.ErrorField(nilEngineErr))
	}

	g.engine.POST(relativePath, func(context *gin.Context) {
		// set token
		context.Set(TOKENKEY, context.GetHeader(TOKENKEY))

		req := reflect.New(ty).Interface()
		if err := context.Bind(req); err != nil {
			log.Logger.Error("Read POST Request Body", log.ErrorField(err))
			context.JSON(200, protocol.PostResponse{
				ErrCode: errCode2,
				Msg:     "read post request body error",
			})
		}

		context.JSON(200, handler(context, req))
	})
}
