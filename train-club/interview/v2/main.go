package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	t := reflect.TypeOf(Author{})       // 获取对象的类型
	for i := 0; i < t.NumField(); i++ { // 获取结构体成员的数量
		name := t.Field(i).Name // Field(i)获取第i个成员的名字
		s, _ := t.FieldByName(name)
		fmt.Println(name, s.Tag)
	}
}
