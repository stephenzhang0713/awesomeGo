package main

import "fmt"

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量hp引用到闭包中
		// 将变量name引用到闭包中, 这个参数在闭包创建时就被固定下来，不会随着后续调用而改变。这就是为什么你可以用generator := playerGen("码神")创建一个生成器，然后调用generator()返回"码神"和150
		return name, hp
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen1() func(string) (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func(name string) (string, int) {
		// 闭包引用的name参数是在调用闭包时传入的
		return name, hp
	}
}

func main() {
	// 创建一个玩家生成器
	generator := playerGen("码神") // "码神"这个字符串被闭包捕获，成为了闭包的环境
	// 返回玩家的名字和血量
	name, hp := generator()
	println(name, hp) // 码神 150

	generator1 := playerGen1()
	name1, hp1 := generator1("码神")
	fmt.Println(name1, hp1) // 码神 150
}
