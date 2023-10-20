package main

func test() {
	l := 1
	a := make([]int, l, l)
	for i := 0; i < l; i++ {
		a[i] = i
	}
}

func main() {
	test()
}
