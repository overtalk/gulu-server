package parse

import (
	"encoding/json"
	"strconv"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func Float(in interface{}) float64 {
	var ret float64
	jsonIn, ok := in.(json.Number)
	if ok {
		in = jsonIn.String()
	}

	switch in.(type) {
	case string:
		inp := in.(string)
		if inp == "" {
			return 0
		}
		var err error
		ret, err = strconv.ParseFloat(inp, 64)
		if err != nil {
			log.Logger.Error("parse to int", log.ErrorField(err), log.Field("input", in))
		}
	case int:
		ret = float64(in.(int))
	case int32:
		ret = float64(in.(int32))
	case int64:
		ret = float64(in.(int64))
	case uint:
		ret = float64(in.(uint))
	case uint32:
		ret = float64(in.(uint32))
	case uint64:
		ret = float64(in.(uint64))
	case float64:
		ret = float64(in.(float64))
	case nil:
		return 0
	default:
		log.Logger.Error("parse to int", log.Field("input", in))
		return 0
	}

	return ret
}
