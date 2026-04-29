package main

import "fmt"

func maxProfit(prices []int) int {
	// 维护一个变量minPrice，记录到当前天为止的最低价格
	// 对于每一天i，计算prices[i] - minPrice作为潜在利润，并更新最大利润
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxResult := 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxResult = max(maxResult, prices[i]-minPrice)

	}

	return maxResult
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
