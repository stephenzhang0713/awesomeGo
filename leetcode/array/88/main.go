package main

import "fmt"

// 过滤器：遍历数组 => 条件 => 过滤后的数组
// 双指针：2个指针用于遍历数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	var res []int
	// 双指针
	var i, j = 0, 0
	// 遍历数组
	for i < m || j < n {
		if j >= n || (i < m && nums1[i] < nums2[j]) {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	for i, num := range res { // 将res数组的值赋值给nums1
		nums1[i] = num
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	fmt.Println(nums2)
}

// i
// 1,2,3,0,0,0
// 2,5,6
// j
