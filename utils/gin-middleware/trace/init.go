package trace

import (
	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set trace ID
		ctx.Set(TraceKey, GenTraceID(Param{PlayerID: "playerID", Url: ctx.FullPath()}))
	}
}
