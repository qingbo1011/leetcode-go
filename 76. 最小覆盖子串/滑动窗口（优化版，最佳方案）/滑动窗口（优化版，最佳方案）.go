package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 使用数组记录字符需求（ASCII 128）
	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	window := [128]int{}
	left, right := 0, 0
	have := 0      // 当前窗口中已满足需求的字符种类数
	needCount := 0 // t中不同字符的种类数
	for _, v := range need {
		if v > 0 {
			needCount++
		}
	}
	minLen := len(s) + 1 // 返回最小子串的长度
	start := 0           // 返回最小子串在s的索引

	for right < len(s) {
		window[s[right]]++
		// 如果窗口内该字符数量刚好达到需求，则have++
		if need[s[right]] == window[s[right]] {
			have++
		}

		// 当所有所需字符都满足时，尝试收缩左边界
		for have == needCount {
			// 更新最小窗口
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			// 移除左边界字符
			// 先判断have是否需要更新：left所在字符是符合覆盖子串条件所need的，则have--
			if need[s[left]] == window[s[left]] {
				have--
			}
			window[s[left]]--
			left++
		}

		right++
	}

	// minLen初始设置len(s)+1，比原始字符串长度+1，如果minLen没有被更新，说明中间没有遇到子串
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

func main() {

}
