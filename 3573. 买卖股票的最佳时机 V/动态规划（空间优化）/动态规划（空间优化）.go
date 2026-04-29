package main

import "math"

func maximumProfit(prices []int, K int) int64 {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// dp[k][state]
	dp := make([][3]int, K+1)

	// 初始化（第 0 天）
	for k := 0; k <= K; k++ {
		dp[k][0] = 0
		dp[k][1] = -prices[0] // 开多
		dp[k][2] = prices[0]  // 开空
	}
	dp[0][1] = math.MinInt32
	dp[0][2] = math.MinInt32

	for i := 1; i < n; i++ {
		p := prices[i]

		// k 必须倒序
		for k := K; k >= 1; k-- {
			// 空仓
			dp[k][0] = max(
				dp[k][0],
				dp[k][1]+p,
				dp[k][2]-p,
			)

			// 多头
			dp[k][1] = max(
				dp[k][1],
				dp[k-1][0]-p,
			)

			// 空头
			dp[k][2] = max(
				dp[k][2],
				dp[k-1][0]+p,
			)
		}

		// k = 0 重置
		dp[0][0] = 0
		dp[0][1] = math.MinInt32
		dp[0][2] = math.MinInt32
	}

	return int64(dp[K][0])
}

func main() {

}
