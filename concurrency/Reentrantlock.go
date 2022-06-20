package main

import (
	"fmt"
	"sync"
)

func foo1(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	foo1(l)
}
