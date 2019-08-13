package parse

import (
	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func Bool(in interface{}) bool {
	var ret bool
	switch in.(type) {
	case bool:
		ret = in.(bool)
	case nil:
		return false
	default:
		log.Logger.Error("parse to int", log.Field("input", in))
		return false
	}

	return ret
}
