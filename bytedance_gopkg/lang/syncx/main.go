package main

import "github.com/bytedance/gopkg/lang/syncx"

var pool = &syncx.Pool{
	New: func() interface{} {
		return &struct{}{}
	},
	NoGC: true,
}

func getPut() {
	var obj = pool.Get().(*struct{})
	defer pool.Put(obj)
}

func main() {
	getPut()
}
