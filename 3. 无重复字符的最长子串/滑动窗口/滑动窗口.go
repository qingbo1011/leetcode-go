package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int, 0)
	res := 0
	left := 0
	for i := range s {
		if index, ok := m[s[i]]; ok { // 字符重复
			left = max(left, index+1) // 新的left值取left当前值和重复元素的索引+1这两者之中的最大值
		}
		m[s[i]] = i // 不管是否更新left，都要更新当前遍历字符的位置！
		res = max(res, i-left+1)
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
