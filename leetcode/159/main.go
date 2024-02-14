package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 用于存储字符出现的次数
	m := make(map[byte]int)
	// 用于存储左右指针
	left, right := 0, 0
	// 用于存储最大长度
	maxLen := 0
	// 用于存储不同字符的个数
	count := 0
	// 右指针右移
	for right < len(s) {
		// 如果当前字符出现的次数为0
		if m[s[right]] == 0 {
			count++
		}
		// 当前字符出现的次数加1
		m[s[right]]++
		// 如果不同字符的个数大于2
		for count > 2 {
			// 如果当前字符出现的次数为1
			if m[s[left]] == 1 {
				count--
			}
			// 当前字符出现的次数减1
			m[s[left]]--
			// 左指针右移
			left++
		}
		// 更新最大长度
		maxLen = max(maxLen, right-left+1)
		// 右指针右移
		right++
	}
	return maxLen
}

func main() {
	s := "eceba"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
