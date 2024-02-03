package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func main() {
	input := []byte(`{"key1":[{},{"key2":{"key3":[1,2,3]}}]}`)
	// no path, returns entire json
	root, err := sonic.Get(input)
	if err != nil {
		fmt.Println(err)
	}
	raw, err := root.Raw() // == string(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(raw)

	// multiple paths
	root, err = sonic.Get(input, "key1", 1, "key2")
	if err != nil {
		fmt.Println(err)
	}

	sub, _ := root.Get("key3").Index(2).Int64() // == 3
	fmt.Println(sub)
}
