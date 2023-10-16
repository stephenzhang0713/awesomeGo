package main

import "fmt"

func hello() {
	fmt.Println("Hello Goroutine!")
}

func main() {
	go hello()
	fmt.Println("main goroutine done!")
}
