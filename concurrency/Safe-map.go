package main

import "sync"

type RWMap struct {
	sync.RWMutex
	m map[int]int
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

func (m *RWMap) Get(k int) (int, bool) { // 从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

func main() {
	rwMap := NewRWMap(10)
	rwMap.Set(10, 10)
	rwMap.Get(10)
}
