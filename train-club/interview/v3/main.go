package main

import "fmt"

type Set map[string]struct{}

func main() {
	set := make(Set)
	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) //3
	if _, ok := set["A"]; ok {
		fmt.Println("A exists")
	}
}
