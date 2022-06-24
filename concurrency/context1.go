package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancelFunc()
	time.Sleep(2 * time.Second)
}
