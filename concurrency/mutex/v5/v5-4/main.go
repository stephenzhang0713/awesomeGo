package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func GoID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}

// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := GoID()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := GoID()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func foo(l RecursiveMutex) {
	fmt.Println("in foo")
	// 调用了l的Lock方法，这会锁定传入的互斥锁。
	// 如果该锁已经被其他goroutine锁定，这个调用将会阻塞，直到锁变得可用。
	l.Lock()
	// defer确保在foo函数的末尾（即函数返回之前）执行Unlock方法。
	// defer通常用于资源清理，这里用于确保互斥锁被释放，以避免死锁
	defer l.Unlock()
	bar(l) // 调用bar函数，并将互斥锁l作为参数传递给它
}

func bar(l RecursiveMutex) {
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
	l := RecursiveMutex{}
	foo(l)
}
