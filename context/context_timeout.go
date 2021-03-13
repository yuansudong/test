package main

import (
	"context"
	"time"
)

type _Request struct {
}

type _Response struct {
}

//
type _ContextWithTimeout struct {
}

func (cwt *_ContextWithTimeout) Do() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cwt.handle(ctx, new(_Request))
	time.Sleep(2 * time.Second)
}

func (cwt *_ContextWithTimeout) handle(ctx context.Context, req *_Request) (*_Response, error) {
	cRspChan := make(chan *_Response)
	go func(cReplyChan chan<- *_Response) {
		time.Sleep(2 * time.Second) // 模拟rpc请求
		cReplyChan <- new(_Response)
	}(cRspChan)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case rsp := <-cRspChan:
		return rsp, nil
	}
}
