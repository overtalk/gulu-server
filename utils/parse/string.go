package parse

import (
	"strconv"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
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
		logger.Logger.Error("parse to string", logger.Field("input", in))
		return ""
	}

	return ret
}
