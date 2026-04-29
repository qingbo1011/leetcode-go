package main

import "fmt"

func canJump(nums []int) bool {
	// dp[i] 表示能否到达下标 i（布尔值）
	dp := make([]bool, len(nums))
	// 初始化：dp[0] = true
	dp[0] = true
	// 状态转移方程：对于每个位置i，如果dp[i]==true，则从i可以跳到i+1到i+nums[i]之间的所有位置，将它们设为true
	for i := 0; i < len(nums); i++ {
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < len(nums); j++ {
				dp[i+j] = true
			}
		}
		// 可以提前检查终点
		if dp[len(nums)-1] {
			return true
		}
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
