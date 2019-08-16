package gate

import (
	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
	. "gitlab.com/SausageShoot/admin-server/utils/trace"
)

func (g *gate) GET(relativePath string, handler module.GETHandler) {
	if g.engine == nil {
		log.Logger.Fatal("add gin get route",
			log.Field("path", relativePath),
			log.Field("handler", handler),
			log.ErrorField(nilEngineErr))
	}

	g.engine.GET(relativePath, func(context *gin.Context) {
		context.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: relativePath}))
		context.String(200, handler(context).Encode())
	})
}
