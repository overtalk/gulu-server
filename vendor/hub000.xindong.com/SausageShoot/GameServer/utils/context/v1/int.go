package context

import (
	"context"
)

func SetContextInt32(ctx context.Context, key interface{}, val int32) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetContextInt32(ctx context.Context, key interface{}) int32 {
	valInterface := ctx.Value(key)
	if val, ok := valInterface.(int32); ok {
		return val
	}
	return 0
}
