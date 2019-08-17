package module

import (
	"context"
	"reflect"
)

type POSTHandler func(ctx context.Context, requestMessage interface{}) interface{}
type GETHandler func(ctx context.Context) interface{}

type Gate interface {
	Start()
	Stop()
	AddStatic(urlPrefix, root string, indexes bool)
	POST(relativePath string, ty reflect.Type, handler POSTHandler)
	GET(relativePath string, handler GETHandler)
}
