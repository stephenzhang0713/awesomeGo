package main

import (
	"fmt"
	"sync"
)

// Mutex 不是可重入的锁

// foo 它接受一个实现了sync.Locker接口的参数l。
// sync.Locker是一个接口，它包含两个方法：Lock和Unlock。
// 任何拥有这两个方法的类型都可以作为sync.Locker接口的实例。
func foo(l sync.Locker) {
	fmt.Println("in foo")
	// 调用了l的Lock方法，这会锁定传入的互斥锁。
	// 如果该锁已经被其他goroutine锁定，这个调用将会阻塞，直到锁变得可用。
	l.Lock()
	// defer确保在foo函数的末尾（即函数返回之前）执行Unlock方法。
	// defer通常用于资源清理，这里用于确保互斥锁被释放，以避免死锁
	defer l.Unlock()
	bar(l) // 调用bar函数，并将互斥锁l作为参数传递给它
}

func bar(l sync.Locker) {
	// bar函数尝试锁定同一个互斥锁l。
	// 但由于这个锁已经在foo函数中被锁定，这个调用会导致死锁，
	// 因为foo函数正在等待bar函数完成，而bar函数又在尝试获得已经被foo函数持有的锁。
	l.Lock()
	// 这里又使用了defer关键字来确保在bar函数结束时调用Unlock方法。
	// 然而，由于前面提到的死锁，这行代码实际上永远不会被执行。
	defer l.Unlock()
	fmt.Println("in bar")
}
func main() {
	// 创建了一个sync.Mutex类型的变量l，并取得它的地址来创建一个互斥锁实例
	l := &sync.Mutex{}
	foo(l)
}
