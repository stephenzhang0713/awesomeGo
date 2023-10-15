package main

import "fmt"

func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {
	s := []int{1, 2, 3}
	res := test("sum: %d", s...)
	println(res)
}
