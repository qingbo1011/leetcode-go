package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	// dp[i]表示nums中找到和为i的排列数（顺序不同算不同）
	dp := make([]int, target+1)
	// dp初始化：和为0的组合有一个：空序列
	dp[0] = 1
	// 状态转移方程：dp[i] = dp[i-nums[0]]+dp[i-nums[1]]+...(条件：i-nums[n]>=0)
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}

	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
