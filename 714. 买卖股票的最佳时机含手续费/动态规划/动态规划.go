package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	// 状态定义：
	//● dp[i][0]：第 i 天结束后，手里没有股票的最大利润。
	//● dp[i][1]：第 i 天结束后，手里持有股票的最大利润。
	dp := make([][2]int, len(prices))
	// 初始化（i=0）：
	//● dp[0][0] = 0：不操作，利润为 0。
	//● dp[0][1] = -prices[0] - fee：第一天买入并支付手续费（相当于将手续费提前到买入时扣除）。
	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee
	// 状态转移（i ≥ 1）：
	//● 不持股：要么昨天也不持股，要么昨天持股并在今天卖出（卖出时不重复扣除手续费，因为买入时已扣）：dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	//● 持股：要么昨天也持股，要么昨天不持股并在今天买入（买入时支付手续费）：dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}

	// 最终答案：返回 max(dp[n-1][0], dp[n-1][1])。由于最后一天持股可能不如不持股，但取最大值也是正确的
	//（实际 dp[n-1][0] 一定不小于 dp[n-1][1]）。
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
