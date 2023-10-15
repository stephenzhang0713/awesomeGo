package main

import (
	"fmt"
	"math"
)

func main() {
	//这里将一个函数当做一个变量一样的操作。
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}
