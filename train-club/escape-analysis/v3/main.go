package main

func test() {
	a := make([]int, 10000, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = i
	}
}

func main() {
	test()
}
