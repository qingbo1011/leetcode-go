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

	// 一维状态定义：dp[j]：表示当前已考虑的数字中，能组成和为j的方案数。
	dp := make([]int, P+1)
	// 初始化：在还没有考虑任何数字时，组成总和为 0 的方案有 1 种（即一个数字都不选）。
	// 其他dp[j]默认为 0，表示无法组成非零总和。
	dp[0] = 1
	// 外层循环：依次考虑每个数字 num。
	for _, num := range nums {
		// 内层循环：j 从 P 递减到 num。
		// dp[j] 更新前代表不考虑当前数字 num 时，组成和 j 的方案数（即上一轮的结果）。
		// dp[j-num] 代表不考虑当前数字 num 时，组成和 j-num 的方案数。
		// 如果我们在当前数字 num 的基础上，加上 j-num 的方案，就得到了包含 num 且总和为 j 的新方案。因此新方案数 = 旧方案数（不选 num）+ 选 num 的方案数。
		// 更新后 dp[j] 就变成了考虑当前数字后，组成和 j 的总方案数。
		for j := P; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}

	return dp[P]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
