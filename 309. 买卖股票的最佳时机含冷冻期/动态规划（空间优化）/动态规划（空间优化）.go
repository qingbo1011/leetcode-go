package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 初始化
	noStockNoCooldown := 0 // dp[i-1][0]
	hold := -prices[0]     // dp[i-1][1]
	noStockCooldown := 0   // dp[i-1][2]
	for i := 1; i < n; i++ {
		newNoStockNoCooldown := max(noStockNoCooldown, noStockCooldown)
		newHold := max(hold, noStockNoCooldown-prices[i])
		newNoStockCooldown := hold + prices[i]
		noStockNoCooldown, hold, noStockCooldown = newNoStockNoCooldown, newHold, newNoStockCooldown
	}
	return max(noStockNoCooldown, noStockCooldown)
}

func main() {

}
