package main

import "fmt"

func main() {
	var a interface{} // a的类型部分和值部分都为nil
	fmt.Printf("a == nil is %t\n", a == nil)

	var b interface{}
	var p *int = nil
	b = p // b的类型部分为*int，值部分为nil
	fmt.Printf("b == nil is %t\n", b == nil)
}
