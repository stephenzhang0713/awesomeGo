package main

import (
	"fmt"
	"github.com/bytedance/gopkg/collection/skipmap"
)

/*
skipmap 是一种基于跳转列表的高性能并发映射。
在典型模式下（一百万次操作，90% LOAD、9% STORE、1% DELETE），skipmap比内置的sync.Map快3~10 倍。
与 sync.Map 不同的是，skipmap 中的项目总是排序的，而且 Load 和 Range 操作是无需等待的（一个 goroutine 只要继续执行操作步骤，就能保证完成操作，与其他 goroutine 的活动无关）。

何时应使用skipmap？
1、需要有序元素
2、并发调用多个操作，如同时使用 Load 和 Store

在这种情况下， sync.Map 更好

	大部分时间只用一个程序访问映射，例如插入一批元素，然后只使用Load（使用内置map更好）
*/
func main() {
	l := skipmap.NewInt()
	for _, v := range []int{10, 12, 15} {
		l.Store(v, v+100)
	}
	v, ok := l.Load(10)
	if ok {
		fmt.Println("skipmap load 10 with value:", v)
	}
	l.Range(func(key int, value interface{}) bool {
		fmt.Println("skipmap range found:", key, value)
		return true
	})
	l.Delete(15)
	fmt.Printf("skipmap contains %d items\n", l.Len())
}
