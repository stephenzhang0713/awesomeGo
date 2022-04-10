map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
## map定义
Go语言中 map的定义语法如下

    map[KeyType]ValueType
其中，

    KeyType:表示键的类型。
    ValueType:表示键对应的值的类型。
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

    make(map[KeyType]ValueType, [cap])
其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
## map基本使用
map中的数据都是成对出现的，map的基本使用示例代码如下：
```
func main() {
    scoreMap := make(map[string]int, 8)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    fmt.Println(scoreMap)
    fmt.Println(scoreMap["小明"])
    fmt.Printf("type of a:%T\n", scoreMap)
}   
```
输出：
```
map[小明:100 张三:90]
100
type of a:map[string]int  
```
map也支持在声明的时候填充元素，例如：
```
func main() {
    userInfo := map[string]string{
        "username": "pprof.cn",
        "password": "123456",
    }
    fmt.Println(userInfo) //
}
```
## 判断某个键是否存在
Go语言中有个判断map中键是否存在的特殊写法，格式如下:

    value, ok := map[key]   
举个例子：
```
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
    v, ok := scoreMap["张三"]
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("查无此人")
    }
} 
```
