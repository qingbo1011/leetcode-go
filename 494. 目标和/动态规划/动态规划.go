package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	// 问题转化：设所有正数和为 P（添加 '+'），负数和为 N（添加 '-'），则：
	// P + N = sum(nums)
	// P - N = target
	// 解得 P = (target + sum) / 2，且 (target + sum) 必须为非负偶数，且 0 ≤ P ≤ sum。
	// 原问题转化为：在nums中选出一个子集，使其和等于P，求方案数（每个数只能选一次，0/1 背包）。
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	// 边界条件
	if target > sum || target < -sum || (target+sum)%2 != 0 {
		return 0
	}
	P := (target + sum) / 2

	// dp数组：dp[i][j]：考虑前i个数字（下标 0 ~ i-1），能组成和为j的方案数（0 ≤ i ≤ n，0 ≤ j ≤ P）
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, P+1)
	}
	// 初始化：
	// dp[0][0] = 1：前0个数字，组成和为0的方案有 1 种（什么都不选）。
	// dp[0][j] = 0（j > 0）：无法组成正数和。
	dp[0][0] = 1
	// 状态转移方程：对于第i个数字 num = nums[i-1]：
	// 不选num：方案数为 dp[i-1][j]
	// 选num（需要 j ≥ num）：方案数为 dp[i-1][j-num]
	// 所以：dp[i][j] = dp[i-1][j] + (j >= num ? dp[i-1][j-num] : 0)
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= P; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][P]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
