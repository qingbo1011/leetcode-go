package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	left, right := 0, len(t)-1
	windowCount := make(map[byte]int) // map记录字符在滑动窗口出现的次数
	needCount := make(map[byte]int)   // 记录目标字符串t中各字符需求次数
	for i := 0; i < len(t); i++ {
		windowCount[s[i]]++
		needCount[t[i]]++
	}
	res := ""
	for right < len(s) {
		if isCovered(needCount, windowCount) {
			if res == "" || len(s[left:right+1]) < len(res) {
				res = s[left : right+1]
			}
			windowCount[s[left]]--
			left++
		} else {
			right++
			if right < len(s) {
				windowCount[s[right]]++
			}
		}
	}

	return res
}

func isCovered(needCount map[byte]int, windowCount map[byte]int) bool {
	for key, value := range needCount {
		// 注意是m中的字符次数小于needCount的情况时，则一定不是子串
		if windowCount[key] < value {
			return false
		}
	}
	return true
}

func main() {

}
