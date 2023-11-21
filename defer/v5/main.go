package main

import "fmt"

func test() {
	defer func() {
		fmt.Println("defer: panic 之前0, 不捕获")
	}()

	defer func() {
		fmt.Println("defer: panic 之前1, 捕获异常")
		// 捕获异常信息
		if err := recover(); err != nil {
			// 输出panic中的错误信息
			fmt.Println(err.(string))
		}
	}()
	// 正常进栈
	defer func() {
		fmt.Println("defer: panic 之前2, 不捕获")
	}()
	//触发defer出栈
	panic("触发异常")
	// 由于在panic之后，不会在执行
	defer func() {
		fmt.Println("defer: panic 之后, 永远执行不到")
	}()
}

func main() {
	test()
	// 由于存在recover捕获panic，main函数流程则正常执行
	fmt.Println("main正常结束")
}
