package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// map也支持在声明的时候填充元素
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 判断某个键是否存在
	v, ok := scoreMap["张三"]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
