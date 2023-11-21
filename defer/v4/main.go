package main

import "fmt"

// 由于返回值在函数定义的时候已经将该变量进行定义
// 在执行return的时候会先执行返回值保存操作
// 而后续的defer函数会改变这个返回值
func test() (i int) { //返回值命名i
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	return i
}

func main() {
	fmt.Println("test:", test())
}
