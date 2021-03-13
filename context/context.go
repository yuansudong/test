package main

import (
	"context"
	"log"
	"time"
)

func main() {
	_ContextWithCancel()
}

func _ContextWithCancel() {
	/*
	 1. 父context的cancel被调用.子context的Done()会收到关闭通知
	 2. 子context的cancel被调用.父context的Done()不会收到关闭通知
	*/

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		aCtx, aCancel := context.WithCancel(ctx)
		aCancel()
		select {
		case <-aCtx.Done():
			log.Println("A协程退出")
		}
	}()
	go func() {
		select {
		case <-ctx.Done():
			log.Println("B协程")
		}
	}()
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
	log.Println(ctx.Err())

}
