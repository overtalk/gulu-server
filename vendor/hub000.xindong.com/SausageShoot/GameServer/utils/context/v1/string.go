package context

import (
	"context"
)

func SetContextString(ctx context.Context, key interface{}, val string) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetContextString(ctx context.Context, key interface{}) string {
	valInterface := ctx.Value(key)
	if val, ok := valInterface.(string); ok {
		return val
	}
	return ""
}
