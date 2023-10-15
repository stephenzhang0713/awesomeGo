package main

import "fmt"

//函数做为一等公民，可以做为参数传递。

func test(fn func() int) int {
	return fn()
}

func fn() int {
	return 200
}

func main() {
	//这是直接使用匿名函数
	s1 := test(func() int {
		return 100
	})
	s1 = test(fn)
	fmt.Println(s1)
}
