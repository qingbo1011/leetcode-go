package main

import "fmt"

func firstMissingPositive(nums []int) int {
	maxNum := nums[0]
	m := make(map[int]struct{})
	for _, v := range nums {
		maxNum = max(maxNum, v)
		if v > 0 {
			m[v] = struct{}{}
		}
	}

	// 如果最大值是小于等于0，那么缺失的第一个正数就是1
	if maxNum <= 0 {
		return 1
	}

	for i := 1; i < maxNum; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return maxNum + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
