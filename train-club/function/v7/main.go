package main

import "fmt"

func FGen(x, y int) (func() int, func(int) int) {
	//求和的匿名函数
	sum := func() int {
		return x + y
	}
	// (x+y) *z 的匿名函数
	avg := func(z int) int {
		return (x + y) * z
	}
	return sum, avg
}

func main() {
	f1, f2 := FGen(1, 2)
	fmt.Println(f1())  //3
	fmt.Println(f2(3)) //9
}
