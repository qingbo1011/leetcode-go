package main

import "fmt"

func maxProfit(prices []int) int {
	// 由于不限交易次数，只要某一天的价格比前一天高，就可以在前一天买入、当天卖出，赚取差价
	// 将所有正的价格差累加起来，就是最大利润
	ans := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			ans = ans + diff
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
