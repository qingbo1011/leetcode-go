package main

import "fmt"

func maxProfit(prices []int) int {
	// 状态定义：定义三种状态（第 i 天结束后）：
	//● dp[i][0]：不持股，且不处于冷冻期（即明天可以自由买入）。
	//● dp[i][1]：持股（持有股票）。
	//● dp[i][2]：不持股，且处于冷冻期（即今天刚卖出，明天不能买入）。
	dp := make([][3]int, len(prices))
	// 初始化（第 0 天）：
	//● dp[0][0] = 0（不买不卖）
	//● dp[0][1] = -prices[0]（买入）
	//● dp[0][2] = 0（不可能在第一天就处于冷冻期，但初始化为0对后续计算无影响，因为状态转移会用max比较）
	dp[0][0], dp[0][1], dp[0][2] = 0, -prices[0], 0
	// 状态转移方程：
	//● dp[i][0]（不持股且非冷冻）：dp[i][0] = max(dp[i-1][0], dp[i-1][2])
	//  ○ 可能来自昨天也不持股且非冷冻（dp[i-1][0]），或者昨天是冷冻期（dp[i-1][2]，今天冷冻期结束）。
	//● dp[i][1]（持股）：dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	//  ○ 可能来自昨天已持股（dp[i-1][1]），或者昨天不持股且非冷冻（dp[i-1][0]）今天买入。
	//● dp[i][2]（不持股且冷冻）：dp[i][2] = dp[i-1][1] + prices[i]
	//  ○ 只能来自昨天持股，今天卖出。
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	// 也可以max(dp[n-1][0], dp[n-1][2])，因为持股状态一定不大于非持股状态
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
