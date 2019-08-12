package parse

import (
	"strconv"

	"hub000.xindong.com/SausageShoot/GameServer/utils/logger/v2"
)

func String(in interface{}) string {
	var ret string
	switch in.(type) {
	case string:
		ret = in.(string)
	case []uint8:
		ret = string(in.([]uint8))
	case int64:
		ret = strconv.FormatInt(in.(int64), 10)
	case int:
		ret = strconv.Itoa(in.(int))
	case float64:
		ret = strconv.FormatFloat(in.(float64), 'E', -1, 64)
	case nil:
		return ""
	default:
		logger.GetLogger().Errorf("parse to string error(unknown) : %v", in)
		return ""
	}

	return ret
}
