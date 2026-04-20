package main

import "fmt"

func coinChange(coins []int, amount int) int {
	// 状态定义：dp[i][j]表示使用前i种硬币（下标 0 ~ i-1）凑成金额j所需的最少硬币个数。
	// 若无法凑成，则值为-1（替代传统的无穷大）。
	// 其中 0 ≤ i ≤ len(coins)，0 ≤ j ≤ amount。
	dp := make([][]int, len(coins)+1)
	// 初始化：
	// dp[0][0] = 0：用 0 种硬币凑 0 元，需要 0 个硬币。
	// dp[0][j] = -1（j > 0）：无法用空集凑出正数金额。
	// 所有dp[i][j]初始为-1（在创建时统一设置）。
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	// 状态转移方程：对于第i种硬币（面值c = coins[i-1]），我们考虑金额j从0到amount：
	// 1.不选当前硬币：若前i-1种硬币能凑成j（即dp[i-1][j] != -1），则dp[i][j]至少为dp[i-1][j]。
	// 2.选一枚当前硬币（前提 j ≥ c 且 dp[i][j-c] != -1）：因为硬币无限，选一枚后剩余金额 j-c仍可用前i种硬币（包括当前面额）继续凑。此时硬币个数 = dp[i][j-c] + 1。
	// (由于硬币无限，选完后还可以继续选，所以是dp[i][j-c]+1，而不是dp[i-1][j-c]+1。这是完全背包与0/1背包的关键区别。)
	// 综合两种情况取最小值（若某情况不可达则忽略）：dp[i][j] = min(不选时的值,选一枚时的值)。实现时需注意 -1 的处理：只有不为 -1 的候选才参与比较。
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if dp[i-1][j] != -1 {
				dp[i][j] = dp[i-1][j]
			}
			if j >= coins[i-1] && dp[i][j-coins[i-1]] != -1 {
				if dp[i-1][j] != -1 {
					dp[i][j] = min(dp[i][j], dp[i][j-coins[i-1]]+1)
				} else {
					dp[i][j] = dp[i][j-coins[i-1]] + 1
				}
			}
		}
	}

	return dp[len(coins)][amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
