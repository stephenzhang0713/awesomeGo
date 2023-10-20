package main

func test() *int {
	// 变量a在函数外部存在引用
	a := 10
	return &a
}

func main() {
	_ = test()
}
