package main

import "fmt"

func change(amount int, coins []int) int {
	// 状态定义：dp[i][j]表示使用前i种硬币（下标 0 ~ i-1）凑成金额j的组合数
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	// 初始化：
	// dp[0][0] = 1：用0种硬币凑0元，有1种组合（空组合）。
	// dp[0][j] = 0（j > 0）：无法凑出正数金额。
	dp[0][0] = 1
	// 状态转移方程：对于第i种硬币（面值c = coins[i-1]），完全背包的组合数递推公式为：dp[i][j] = dp[i-1][j] + dp[i][j-c]   (当 j ≥ c)
	// 不选第i种硬币：组合数 = dp[i-1][j]
	// 选一枚第i种硬币：由于硬币无限，选一枚后剩余金额j-c仍可用前i种硬币（包括当前面额）继续凑，组合数 = dp[i][j-c]
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i-1] {
				dp[i][j] = dp[i][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
