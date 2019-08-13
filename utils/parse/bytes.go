package parse

import (
	"gitlab.com/SausageShoot/admin-server/utils/log"
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
		log.Logger.Error("parse to string", log.Field("input", in))
		return nil
	}

	return ret
}
