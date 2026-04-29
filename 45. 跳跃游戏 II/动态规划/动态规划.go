package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	// dp[i] 表示从起点跳到下标 i 所需的最小跳跃次数。
	dp := make([]int, len(nums))
	// 初始化：
	// dp[0] = 0（起点不需要跳跃）
	// dp[i] = 一个很大的数（如 math.MaxInt） 表示尚未到达，i = 1..n-1
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	//状态转移：遍历每个位置i（从 0 到 n-1），对于当前位置i，已知 p[i]是最小步数。
	//然后从i出发，可以跳1到nums[i]步，到达位置 i+j（1 ≤ j ≤ nums[i] 且 i+j < n）。
	//那么到达i+j的步数可以是dp[i]+1我们取最小值：dp[i+j] = min(dp[i+j], dp[i] + 1)
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}

	return dp[len(nums)-1]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
