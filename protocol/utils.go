package protocol

import (
	"encoding/json"
	"reflect"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

func GetRequest(ty reflect.Type, data []byte) interface{} {
	req := reflect.New(ty).Interface()
	if err := json.Unmarshal(data, req); err != nil {
		log.Logger.Error("Unmarshal Http Request Data",
			log.Field("request type", ty),
			log.Field("data", string(data)))
		return nil
	}
	return req
}
