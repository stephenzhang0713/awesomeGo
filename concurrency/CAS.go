package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	a := int32(10)
	ok := atomic.CompareAndSwapInt32(&a, 10, 100)
	fmt.Println(a, ok)
	ok = atomic.CompareAndSwapInt32(&a, 10, 50)
	fmt.Println(a, ok)
}
