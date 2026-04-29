package main

import "fmt"

func maxProfit(prices []int) int {
	// 状态定义：
	// dp[i][0] 表示第 i 天结束后，手里没有股票时的最大利润。
	// dp[i][1] 表示第 i 天结束后，手里有股票时的最大利润。
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	// 初始化：
	// dp[0][0] = 0（第一天没买）
	// dp[0][1] = -prices[0]（第一天买入，利润为负）
	dp[0][0], dp[0][1] = 0, -prices[0]
	// 状态转移：
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]) 。今天没持有：要么之前就没有，要么之前有但今天卖出了
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) 。今天持有：要么之前就有，要么今天买入
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
