package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	go dosomething(100, &wg)
	go dosomething(110, &wg)
	go dosomething(120, &wg)
	go dosomething(130, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)

		fmt.Println("后台执行, duration:", duration)
		wg.Done()
	}()
}
