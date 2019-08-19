package module

import (
	"context"
)

type Auth interface {
	CurrentPlayer(ctx context.Context) interface{}
	Login(ctx context.Context, requestMessage interface{}) interface{}
}
