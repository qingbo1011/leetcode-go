package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	ans := ""
	lengthS, lengthT := len(s), len(t)
	if lengthS < lengthT {
		return ""
	}
	// 得到t字符串对应的数组arrT（方便和滑动窗口类的字符串进行比较，判断滑动窗口是否合法）
	tArr := [58]int{}
	for _, c := range t {
		tArr[c-'A']++
	}
	// 准备维护滑动窗口
	left, i := 0, 0
	windowArr := [58]int{}
	for i < lengthS {
		windowArr[s[i]-'A']++
		if check(windowArr, tArr) { // 如果符合条件了就要收缩滑动窗口以得到最小的覆盖子串
			for check(windowArr, tArr) {
				windowArr[s[left]-'A']--
				left++
			}
			ans = min(ans, s[left-1:i+1])
		}
		i++
	}
	return ans
}

func check(windowArr [58]int, tArr [58]int) bool {
	for i := 0; i < 58; i++ {
		if windowArr[i] < tArr[i] {
			return false
		}
	}
	return true
}

func min(ans string, s string) string {
	if ans == "" || len(s) < len(ans) {
		return s
	}
	return ans
}
