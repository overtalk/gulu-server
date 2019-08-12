package parse

import (
	"errors"

	"hub000.xindong.com/SausageShoot/GameServer/utils/logger/v2"
)

func Bool(in interface{}) bool {
	var ret bool
	switch in.(type) {
	case bool:
		ret = in.(bool)
	case nil:
		return false
	default:
		logger.GetLogger().Error(errors.New("parse to int error(unknown) : error type"))
		return false
	}

	return ret
}
