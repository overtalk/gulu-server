package module

import (
	"reflect"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

type Handler func(requestMessage interface{}) protocol.Response

type Gate interface {
	Start()
	Stop()
	AddHandler(httpMethod, relativePath string, ty reflect.Type, handler Handler)
}
