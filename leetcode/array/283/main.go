package main

import "fmt"

// 过滤器
// filter(nums) => 条件 => 新nums
// 遍历
func moveZeroes(nums []int) {
	var cnt = 0                // 计数指针
	for _, num := range nums { // 遍历
		if num != 0 { // 过滤
			nums[cnt] = num // 条件非0的元素，放到前面的位置
			cnt++           //并计数+1
		}
	}
	for cnt < len(nums) { // 剩下的位置补0
		nums[cnt] = 0 // 从cnt开始，后面的位置都补0
		cnt++         //并计数+1
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
