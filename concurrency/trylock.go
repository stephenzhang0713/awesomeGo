package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}

// Mutex 扩展一个Mutex结构
type Mutex struct {
	sync.Mutex
}

func (m *Mutex) Count() int { // 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v>>mutexWaiterShift + (v & mutexLocked)
	return int(v)
}

// IsLocked 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// IsWoken 是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// IsStarving 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

func count() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func try() {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock() // 尝试获取到锁
	if ok {            // 获取成功
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}

	// 没有获取到
	fmt.Println("can't get the lock")
}

func main() {
	count()
}
