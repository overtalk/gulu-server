package parse

import (
	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func Bytes(in interface{}) []byte {
	var ret []byte
	switch in.(type) {
	case []byte:
		ret = in.([]byte)
	case string:
		ret = []byte(in.(string))
	case nil:
		return nil
	default:
		logger.Logger.Error("parse to string", logger.Field("input", in))
		return nil
	}

	return ret
}
