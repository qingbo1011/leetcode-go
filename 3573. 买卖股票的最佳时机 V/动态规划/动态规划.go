package main

import (
	"fmt"
	"math"
)

func maximumProfit(prices []int, K int) int64 {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 状态定义：dp[i][k][state]
	//● i：第 i 天
	//● k：最多还能做 k 笔交易
	//●  state：
	//  ○ 0：空仓
	//  ○ 1：持有多头（说明正在做普通交易（已经买了，还没卖））
	//  ○ 2：持有空头（说明正在做空交易（已经卖了，还没买回））
	dp := make([][][3]int, n)
	for i := range dp {
		dp[i] = make([][3]int, K+1)
	}

	// 初始化：第 0 天：
	//dp[0][k][0] = 0
	//dp[0][k][1] = -prices[0]   // 开多（买入）
	//dp[0][k][2] = prices[0]    // 开空（卖出）
	//但必须修正非法状态：
	//dp[i][0][1] = -∞	// k=0不能做交易了
	//dp[i][0][2] = -∞	// k=0不能做交易了
	//dp[i][0][0] = 0
	for k := 0; k <= K; k++ {
		dp[0][k][0] = 0
		dp[0][k][1] = -prices[0]
		dp[0][k][2] = prices[0]
	}
	dp[0][0][1] = math.MinInt32
	dp[0][0][2] = math.MinInt32

	// 状态转移：
	//1️⃣ 空仓：
	//dp[i][k][0] = max(
	//    dp[i-1][k][0],         // 什么都不做
	//    dp[i-1][k][1] + p,     // 做多（卖出）
	//    dp[i-1][k][2] - p      // 做空（买回）
	//)
	//2️⃣ 持有多头（说明正在做普通交易（已经买了，还没卖））：
	//dp[i][k][1] = max(
	//    dp[i-1][k][1],         // 继续持有
	//    dp[i-1][k-1][0] - p    // 开多（买入，消耗一次交易）
	//)
	//3️⃣ 持有空头（说明正在做空交易（已经卖了，还没买回））：
	//dp[i][k][2] = max(
	//    dp[i-1][k][2],         // 继续持有空头
	//    dp[i-1][k-1][0] + p    // 开空（卖出，消耗一次交易）
	//)
	for i := 1; i < n; i++ {
		// k = 0 特殊处理
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt32
		dp[i][0][2] = math.MinInt32
		for k := 1; k <= K; k++ {
			p := prices[i]
			// 空仓
			dp[i][k][0] = max(
				dp[i-1][k][0],
				dp[i-1][k][1]+p,
				dp[i-1][k][2]-p,
			)
			// 持有多头（说明正在做普通交易（已经买了，还没卖））
			dp[i][k][1] = max(
				dp[i-1][k][1],
				dp[i-1][k-1][0]-p,
			)
			// 持有空头（说明正在做空交易（已经卖了，还没买回））
			dp[i][k][2] = max(
				dp[i-1][k][2],
				dp[i-1][k-1][0]+p,
			)
		}
	}

	//最终答案：dp[n-1][k][0]（必须空仓才算完成交易）
	return int64(dp[n-1][K][0])
}

func main() {
	fmt.Println(maximumProfit([]int{1, 7, 9, 8, 2}, 2))
}
