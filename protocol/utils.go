package protocol

import (
	"encoding/json"
	"reflect"

	"gitlab.com/SausageShoot/admin-server/utils/logger"
)

func GetRequest(ty reflect.Type, data []byte) interface{} {
	req := reflect.New(ty).Interface()
	if err := json.Unmarshal(data, req); err != nil {
		logger.Logger.Error("Unmarshal Http Request Data",
			logger.Field("request type", ty),
			logger.Field("data", string(data)))
		return nil
	}
	return req
}
