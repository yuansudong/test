package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cExit := make(chan struct{})
	wg.Add(3)
	go func() {
		defer wg.Done()
		select {
		case <-cExit:
			fmt.Println("g1 退出")
			return
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-cExit:
			fmt.Println("g2 退出")
			return
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-cExit:
			fmt.Println("g3 退出")
			return
		}
	}()
	close(cExit)
	wg.Wait()
}
