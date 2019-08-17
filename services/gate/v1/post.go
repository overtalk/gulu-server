package gate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"

	"gitlab.com/SausageShoot/admin-server/module"
	"gitlab.com/SausageShoot/admin-server/protocol"
	"gitlab.com/SausageShoot/admin-server/utils/log"
	. "gitlab.com/SausageShoot/admin-server/utils/trace"
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
		// set trace ID
		context.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: relativePath}))
		// set token
		context.Set(TOKENKEY, context.GetHeader(TOKENKEY))

		fmt.Println("token = ", context.GetHeader(TOKENKEY))

		body, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			log.Logger.Fatal("Read POST Request Body", log.ErrorField(err))
			context.String(200, protocol.PostResponse{
				ErrCode: errCode2,
				Msg:     "read post request body error",
			}.Encode())
		}

		resp := handler(context, protocol.GetRequest(ty, body))
		data, err := json.Marshal(resp)
		if err != nil {
			log.Logger.Error("Response Encode",
				log.ErrorField(err),
				log.Field("method", "POST"),
				log.Field("path", relativePath),
				log.Field("resp", resp))
			data = []byte("")
		}
		context.String(200, string(data))
	})
}
