package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

var token = rand.Int63()

// TokenRecursiveMutex Token方式的递归锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

// Lock 请求锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token { //如果传入的token和持有锁的token一致，说明是递归调用
		m.recursion++
		return
	}
	m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// Unlock 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 释放其它token持有的锁
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--         // 当前持有这个锁的token释放锁
	if m.recursion != 0 { // 还没有回退到最初的递归调用
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
	m.Mutex.Unlock()
}

func foo(l *TokenRecursiveMutex) {
	fmt.Println("in foo")
	// 调用了l的Lock方法，这会锁定传入的互斥锁。
	// 如果该锁已经被其他goroutine锁定，这个调用将会阻塞，直到锁变得可用。
	l.Lock(token)
	// defer确保在foo函数的末尾（即函数返回之前）执行Unlock方法。
	// defer通常用于资源清理，这里用于确保互斥锁被释放，以避免死锁
	defer l.Unlock(token)
	bar(l) // 调用bar函数，并将互斥锁l作为参数传递给它
}

func bar(l *TokenRecursiveMutex) {
	// bar函数尝试锁定同一个互斥锁l。
	// 但由于这个锁已经在foo函数中被锁定，这个调用会导致死锁，
	// 因为foo函数正在等待bar函数完成，而bar函数又在尝试获得已经被foo函数持有的锁。
	l.Lock(token)
	// 这里又使用了defer关键字来确保在bar函数结束时调用Unlock方法。
	// 然而，由于前面提到的死锁，这行代码实际上永远不会被执行。
	defer l.Unlock(token)
	fmt.Println("in bar")
}
func main() {
	// 创建了一个sync.Mutex类型的变量l，并取得它的地址来创建一个互斥锁实例
	l := TokenRecursiveMutex{}
	foo(&l)
}
