package main

import (
	"fmt"
	"github.com/bytedance/gopkg/collection/zset"
)

/**
zset 提供了并发安全排序集，可用作 Redis zset 的本地替代。
与其他集合的主要区别在于，集合的每个值都与一个分数相关联，该分数用于将排序后的集合按分数从小到大排序。
在进行 Add(ZADD) 和 Remove(ZREM) 操作时，zset 的时间复杂度为 O(log(N)) ；在进行 Contains 操作时，时间复杂度为 O(1) 。

特点
	并发安全 API
	数值按得分排序
	相当于 redis 的实现
	快速skiplist级别随机化
*/

func main() {
	z := zset.NewFloat64()
	values := []string{"Alice", "Bob", "Zhang"}
	scores := []float64{90, 89, 59.9}
	for i := range values {
		z.Add(scores[i], values[i])
	}
	s, ok := z.Score("Alice")
	if ok {
		fmt.Println("Alice's score is", s)
	}
	n := z.Count(0, 60)
	fmt.Println("There are", n, "people whose score is between 0 and 60")

	for _, n := range z.Range(0, -1) {
		fmt.Println("zset range found", n.Value, n.Score)
	}
}
