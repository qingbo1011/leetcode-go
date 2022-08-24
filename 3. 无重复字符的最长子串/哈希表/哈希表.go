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
	ans := 0
	m := make(map[byte]int, len(s)) // key为字符，value为该字符在字符串s中的索引
	tmp := 0
	for i := 0; i < len(s); i++ {
		if value, ok := m[s[i]]; ok { // map中已存在
			m = make(map[byte]int, len(s)) // 直接清空map
			ans = max(ans, tmp)
			tmp = 0
			i = value
		} else { // map中不存在
			m[s[i]] = i
			tmp++
		}
	}
	return max(ans, tmp)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
