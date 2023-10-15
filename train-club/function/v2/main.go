package main

import "fmt"

//在将函数做为参数的时候，我们可以使用类型定义，将函数定义为类型，这样便于阅读

// FormatFunc 定义函数类型。
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func formatFunc(s string, x, y int) string {
	return fmt.Sprintf(s, x, y)
}

func main() {
	s2 := format(formatFunc, "%d, %d", 10, 20)
	fmt.Println(s2)
}
