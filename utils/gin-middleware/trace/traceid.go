package trace

import (
	"fmt"
	"time"
)

const TraceKey = "traceId"

type Param struct {
	PlayerID interface{}
	Url      string
}

func GenTraceID(p Param) string {
	now := time.Now()
	return fmt.Sprintf("%s_%v_%s_%d_%d",
		now.String()[0:19], p.PlayerID, p.Url, now.Unix(), now.Nanosecond())
}
