package main

import (
	"context"
	"log"
)

type _ContextWithValue struct {
	key string
	val string
}

func _NewContextWithValue() *_ContextWithValue {

	return &_ContextWithValue{
		key: "hello",
		val: "world",
	}
}

func (cwv *_ContextWithValue) Do() {
	cwv.Func1(context.Background())
}

func (cwv *_ContextWithValue) Func1(ctx context.Context) *_ContextWithValue {

	cwv.Func2(context.WithValue(ctx, cwv.key, cwv.val))
	return cwv
}

// Func2 用于获得func1设置的值
func (cwv *_ContextWithValue) Func2(ctx context.Context) *_ContextWithValue {
	log.Println(ctx.Value(cwv.key))
	return cwv
}
