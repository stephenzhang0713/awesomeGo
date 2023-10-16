package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("发送成功")
}
