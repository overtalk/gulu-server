package parse

import (
	"encoding/json"
	"errors"
	"strconv"

	"hub000.xindong.com/SausageShoot/GameServer/utils/logger/v2"
)

func Int(in interface{}) int64 {
	var ret int64
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
		ret, err = strconv.ParseInt(inp, 10, 64)
		if err != nil {
			logger.GetLogger().Errorf("parse to int error(string) : %v", err)
		}
	case int:
		ret = int64(in.(int))
	case int32:
		ret = int64(in.(int32))
	case int64:
		ret = in.(int64)
	case uint:
		ret = int64(in.(uint))
	case uint32:
		ret = int64(in.(uint32))
	case uint64:
		ret = int64(in.(uint64))
	case float64:
		ret = int64(in.(float64))
	case nil:
		return 0
	default:
		logger.GetLogger().Error(errors.New("parse to int error(unknown) : error type"))
		return 0
	}

	return ret
}
