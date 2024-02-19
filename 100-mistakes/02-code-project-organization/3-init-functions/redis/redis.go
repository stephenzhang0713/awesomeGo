package redis

import "fmt"

func init() {
	fmt.Println("init")
}

func Store(key, value string) error {
	return nil
}
