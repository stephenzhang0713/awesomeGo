package main

import (
	"fmt"
	"sync"
)

func foo2(l *sync.RWMutex) {
	fmt.Println("in foo")
	l.Lock()
	bar1(l)
	l.Unlock()
}

func bar1(l *sync.RWMutex) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.RWMutex{}
	foo2(l)
}
