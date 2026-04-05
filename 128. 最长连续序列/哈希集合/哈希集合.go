package main

import "fmt"

func longestConsecutive(nums []int) int {
	res := 0

	// 用map模拟集合，去重并便于 O(1) 查找
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	// 注意这里只需遍历set，因为set已经帮我们去重了（遍历nums会超时！）
	for num := range set {
		// 只有当num-1不存在时，才以num为起点开始查找(从最小值开始，找最长连续序列)
		if !set[num-1] {
			currentNum := num
			currentLength := 1
			for set[currentNum+1] {
				currentNum++
				currentLength++
			}
			if currentLength > res {
				res = currentLength
			}
		}
	}

	return res
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
