package main

func maxProfit(prices []int, fee int) int {
	buy := prices[0] // 当前持有的买入成本（已考虑手续费）
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			// 找到更便宜的买入点，更新成本
			buy = prices[i]
		} else if prices[i] > buy+fee {
			// 当前价格卖出有利润
			profit = profit + prices[i] - buy - fee
			// 卖出后，可以立即以当前价格买入（但扣除手续费后才相当于成本）,这样如果后面继续上涨，就可以继续累加利润
			// 为什么 buy 要更新为 prices[i] - fee？
			// 因为我们在 prices[i] 处卖出了，但之后价格若继续上涨，我们实际上不应该只赚取一段，而应该连续交易。
			buy = prices[i] - fee
		}
	}

	return profit
}

func main() {

}
