package main

import "fmt"

func change(amount int, coins []int) int {
	// 一维状态定义：dp[j]表示使用当前已考虑过的硬币凑成金额j的组合数。
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := 0; j <= amount; j++ {
			dp[j] = dp[j]
			if j >= coin {
				dp[j] = dp[j] + dp[j-coin]
			}
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
