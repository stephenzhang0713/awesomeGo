package main

import "fmt"

func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	// 执行Return语句后，Go会创建一个临时变量保存返回值
	return i
}

func test1() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	// 对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存
	return i
}

func main() {
	//fmt.Println("return", test()) // test返回值没有被修改
	fmt.Println("return", test1()) // test返回值被修改
}

// defer2
// defer1
// return 0
