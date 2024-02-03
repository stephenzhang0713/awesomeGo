package main

import (
	"fmt"
	"net/http"
	"time"
)

func getFromBaidu() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get from baidu failed, err:", err)
		return
	}
	fmt.Println(resp.Status)
}

func add(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y can not be zero")
	}
	return x / y, nil
}

func main() {
	// 并发输出
	go func() {
		fmt.Println("Hello from a goroutine")
	}()

	// 并发获取百度首页
	go getFromBaidu()

	go add(1, 2)

	time.Sleep(time.Second)
}
