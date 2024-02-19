package main

import (
	"100-mistakes/02-code-project-organization/3-init-functions/redis"
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
