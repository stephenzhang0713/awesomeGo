package main

import "fmt"

// 有序
// 过滤器 => filter(nums) => 条件 => 新nums
func removeDuplicates(nums []int) int {
	var cnt = 0
	for i, num := range nums {
		// 第一个元素或者当前元素不等于前一个元素
		// 保留当前元素 => nums[cnt] = num
		if i == 0 || nums[i] != nums[i-1] {
			nums[cnt] = num
			cnt++
		}
	}
	return cnt
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
