package main

import (
	"fmt"
	"sync"
)

// Lock/Unlock 不是成对出现

func foo() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world")
}

func main() {
	foo()
}
