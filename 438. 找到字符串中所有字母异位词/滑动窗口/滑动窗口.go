package main

import "fmt"

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	var res []int
	lengthS := len(s)
	lengthP := len(p)
	if lengthS < lengthP {
		return res
	}
	var arrP [123]int
	for i := 0; i < lengthP; i++ {
		arrP[p[i]]++
	}
	// 双指针就位
	var arr [123]int
	left, right := 0, 0
	for right = 0; right < lengthP; right++ {
		arr[s[right]]++
	}
	// 维护滑动窗口
	right-- // right在上一个循环中退出循环后值为lengthP，所以这里需要right--
	for right < lengthS {
		if arr == arrP {
			res = append(res, left)
		}
		left++
		right++
		arr[s[left-1]]--
		if right < lengthS { // 防止right到最后一位时，s[right]溢出的情况
			arr[s[right]]++
		}
	}
	return res
}
