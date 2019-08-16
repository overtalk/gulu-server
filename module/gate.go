package module

import (
	"context"
	"reflect"

	"gitlab.com/SausageShoot/admin-server/protocol"
)

type POSTHandler func(ctx context.Context, requestMessage interface{}) protocol.Response
type GETHandler func(ctx context.Context) protocol.Response

type Gate interface {
	Start()
	Stop()
	AddStatic(urlPrefix, root string, indexes bool)
	POST(relativePath string, ty reflect.Type, handler POSTHandler)
	GET(relativePath string, handler GETHandler)
}
