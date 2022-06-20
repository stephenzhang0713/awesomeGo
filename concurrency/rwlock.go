package main

import (
	"sync"
	"time"
)

// Counter1 一个线程安全的计数器
type Counter1 struct {
	mu    sync.RWMutex
	count uint64
}

// Incr 使用写锁保护
func (c *Counter1) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 使用读锁保护
func (c *Counter1) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var counter Counter1
	for i := 0; i < 10; i++ {
		go func() {
			for {
				counter.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}
