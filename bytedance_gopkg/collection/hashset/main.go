package main

import (
	"fmt"
	"github.com/bytedance/gopkg/collection/hashset"
)

/*
在此存储库中，我们实现了一种基础数据结构：Set 基于 golang 中的 Map。我们有
Add(value int64) 将指定元素添加到此集合中。
Contains(value int64) bool 如果该集合包含指定元素，则返回 true。
Remove(value int64) 从集合中删除指定元素。
Range(f func(value int64) bool) 注：函数 f 以集合中的元素为参数依次执行，直到 f 返回 false。
Len() int 返回该集合的元素个数。

Hashset 不能保证并发安全。如果确实需要并发安全集，请使用 skipset
*/
func main() {
	l := hashset.NewInt()
	for _, v := range []int{10, 12, 15, 15} {
		if l.Add(v) {
			fmt.Println("hashset add", v)
		}
	}
	l.Range(func(value int) bool {
		fmt.Println("hashset range found ", value)
		return true
	})
	//if l.Contains(10) {
	//	fmt.Println("hashset contains 10")
	//}
	//l.Range(func(value int) bool {
	//	fmt.Println("hashset range found ", value)
	//	return true
	//})
	//l.Remove(15)
	//fmt.Printf("hashset contains %d\n", l.Len())
}
