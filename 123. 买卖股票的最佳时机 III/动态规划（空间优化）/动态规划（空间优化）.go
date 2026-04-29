package main

func maxProfit(prices []int) int {
	// 由于最多完成两笔交易，我们可以维护四个状态：
	// buy1：第一次买入后的利润（负数）
	// sell1：第一次卖出后的利润
	// buy2：第二次买入后的利润
	// sell2：第二次卖出后的利润
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}

	return sell2
}

func main() {

}
