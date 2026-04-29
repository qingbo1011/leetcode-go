package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// 状态定义：
	// dp[i][k][0] 表示：第 i 天结束后，最多还能做 k 次交易，且手里没有股票时的最大利润。
	// dp[i][k][1] 表示：第 i 天结束后，最多还能做 k 次交易，且手里持有股票时的最大利润。
	dp := make([][3][2]int, len(prices))
	// 初始化：
	// dp[0][k][0] = 0（第一天不持股）
	// dp[0][k][1] = -prices[0]（第一天买入）
	// dp[i][0][1] = -∞（不允许交易却持股，是不可能状态）
	// dp[i][0][0] = 0
	for k := 0; k <= 2; k++ {
		dp[0][k][0] = 0
		dp[0][k][1] = -prices[0]
	}
	dp[0][0][1] = math.MinInt32 // 不合法状态
	// 状态转移：
	// 今天不持股：dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// (要么昨天就没持股,要么昨天持股，今天卖出)
	// 今天持股：dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	// (要么昨天就持股,要么昨天没持股，今天买入（消耗一次交易次数）)
	for i := 1; i < len(prices); i++ {
		// k=0 特殊处理
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt32
		for k := 1; k <= 2; k++ {
			// 不持股
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 持股
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	// 因为交易次数越多利润不会更少（可以不交易），所以直接取dp[n-1][2][0]也可
	return max(dp[len(prices)-1][0][0], dp[len(prices)-1][1][0], dp[len(prices)-1][2][0])
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
