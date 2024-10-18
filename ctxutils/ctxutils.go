package ctxutils

import (
	"context"
)

type ContextKeyString struct {
	Key string
}

type ContextKeyInt struct {
	Key int
}

func SetStringValue(ctx context.Context, key ContextKeyString, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetStringValue(ctx context.Context, key ContextKeyString) (string, bool) {
	value, ok := ctx.Value(key).(string)
	return value, ok
}

func SetIntValue(ctx context.Context, key ContextKeyInt, value int) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetIntValue(ctx context.Context, key ContextKeyInt) (int, bool) {
	value, ok := ctx.Value(key).(int)
	return value, ok
}
