package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i { // 如果最大距离（所在下标）都小于或等于i了，说明i之前的数字无论如何也无法跳跃到i之后（中途遇到0卡住了）
			return false
		}
	}
	return farthest >= n-1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
