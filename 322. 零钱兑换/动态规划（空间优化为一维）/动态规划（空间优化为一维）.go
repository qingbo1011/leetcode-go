package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	// 一维状态定义：dp[j]表示凑成金额j所需的最少硬币个数，-1表示不可达。
	dp := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		dp[j] = -1
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := 0; j <= amount; j++ {
			if dp[j] != -1 {
				dp[j] = dp[j]
			}
			if j >= coin && dp[j-coin] != -1 {
				if dp[j] != -1 {
					dp[j] = min(dp[j], dp[j-coin]+1)
				} else {
					dp[j] = dp[j-coin] + 1
				}
			}
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
