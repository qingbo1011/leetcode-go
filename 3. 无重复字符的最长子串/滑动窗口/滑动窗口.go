package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int) // 记录字符上次出现的索引
	left := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			left = max(left, index+1) // 出现重复字符左指针右移，不能回退
		}
		m[s[i]] = i // 更新字符最新位置
		res = max(res, i-left+1)
	}

	return res
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
