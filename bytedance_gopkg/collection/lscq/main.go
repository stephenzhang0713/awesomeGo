package main

import (
	"fmt"
	"github.com/bytedance/gopkg/collection/lscq"
)

/**
LSCQ 是 Go 语言中的一种可扩展、无界、多生产者和多消费者 FIFO 队列。

在基准测试（AMD 3700x，运行频率 3.6 GHZ，-cpu=16）中，LSCQ 在大多数情况下都比基于锁的链接队列高出 5 至 6 倍。
由于内置通道是有界队列，我们只能在 EnqueueDequeuePair 中进行比较，在这种情况下，LSCQ 的性能是内置通道的 8~9 倍。
*/

func main() {
	q := lscq.NewUint64()
	q.Enqueue(100)
	q.Enqueue(110)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
