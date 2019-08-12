package parse

import (
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func Bool(in interface{}) bool {
	var ret bool
	switch in.(type) {
	case bool:
		ret = in.(bool)
	case nil:
		return false
	default:
		logger.Logger.Error("parse to int", logger.Field("input", in))
		return false
	}

	return ret
}
