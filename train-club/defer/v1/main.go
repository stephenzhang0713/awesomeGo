package main

import "fmt"

func main() {
	i := 0
	defer func(a int) {
		fmt.Println(a)
	}(i)
	i++
}
