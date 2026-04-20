package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	count := 0
	backtrack(nums, target, 0, 0, &count)
	return count
}

func backtrack(nums []int, target int, index int, sum int, count *int) {
	if index == len(nums) {
		if sum == target {
			*count++
		}
		return
	}
	// 选择加号
	backtrack(nums, target, index+1, sum+nums[index], count)
	// 选择减号
	backtrack(nums, target, index+1, sum-nums[index], count)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
