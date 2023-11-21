package main

import "fmt"

func deferFunc() {
	fmt.Println("defer func called")
}

func returnFunc() int {
	fmt.Println("return func called")
	return 1
}

func returnAndDefer() int {
	// 后执行
	defer deferFunc()
	// 先执行
	return returnFunc()
}

func main() {
	returnAndDefer()
}
