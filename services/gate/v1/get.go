package gate

import (
	"encoding/json"
	"fmt"

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
		// set trace id
		context.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: relativePath}))
		// set token
		context.Set(TOKENKEY, context.GetHeader(TOKENKEY))

		fmt.Println("token = ", context.GetHeader(TOKENKEY))

		resp := handler(context)
		data, err := json.Marshal(resp)
		if err != nil {
			log.Logger.Error("Response Encode",
				log.ErrorField(err),
				log.Field("method", "GET"),
				log.Field("path", relativePath),
				log.Field("resp", resp))
			data = []byte("")
		}
		context.String(200, string(data))
	})
}
