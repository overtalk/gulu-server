package gate

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func (g *gate) GET(relativePath string, handler module.GETHandler) {
	if g.engine == nil {
		log.Logger.Fatal("add gin get route",
			log.Field("path", relativePath),
			log.Field("handler", handler),
			log.ErrorField(nilEngineErr))
	}

	g.api.GET(relativePath, func(context *gin.Context) {
		context.IndentedJSON(200, handler(context))
	})
}
