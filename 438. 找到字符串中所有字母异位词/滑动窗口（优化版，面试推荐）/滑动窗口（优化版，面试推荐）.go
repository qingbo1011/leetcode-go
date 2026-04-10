package main

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	// 初始化
	pCount := [26]int{}
	windowCount := [26]int{}
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		windowCount[s[i]-'a']++
	}
	// 如果相等先记录
	if pCount == windowCount {
		res = append(res, 0)
	}

	for left := 1; left < len(s)-len(p)+1; left++ {
		right := left + len(p) - 1
		// 优化：移除windowCount左边的元素，添加新增的元素
		windowCount[s[left-1]-'a']--
		windowCount[s[right]-'a']++
		if windowCount == pCount {
			res = append(res, left)
		}
	}

	return res
}

func main() {

}
