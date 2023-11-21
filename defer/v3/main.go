package main

import "fmt"

// return语句->defer2->defer1->返回值
// defer2先于defer1执行

func test() int {
	var i int
	defer func() {
		i++
		//作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=2
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		//作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=1
		fmt.Println("defer2", i)
	}()
	//先执行return i, 把i的值给到一个临时变量，作为函数返回值
	return i
}

func main() {
	// defer 和 return之间的顺序是先返回值, i=0，后defer
	fmt.Println(test())
}
