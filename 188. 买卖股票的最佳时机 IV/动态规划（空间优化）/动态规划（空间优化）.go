package main

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}
	// 特判：当 k >= n/2 时，相当于无限次交易，贪心即可
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}
	// cash[j] 表示最多还能进行 j 次交易且不持股的最大利润
	// hold[j] 表示最多还能进行 j 次交易且持股的最大利润
	cash := make([]int, k+1)
	hold := make([]int, k+1)
	// 初始化第0天
	for j := 0; j <= k; j++ {
		cash[j] = 0
		hold[j] = -prices[0]
	}
	hold[0] = math.MinInt32 // j=0 不能持股

	for i := 1; i < n; i++ {
		p := prices[i]
		// 从大到小更新 j，避免覆盖需要的前一天状态
		for j := k; j >= 1; j-- {
			cash[j] = max(cash[j], hold[j]+p)
			hold[j] = max(hold[j], cash[j-1]-p)
		}
		// j=0 单独处理
		cash[0] = max(cash[0], hold[0]+p)
		hold[0] = math.MinInt32 // 保持不可持股
	}
	// 答案即为最多还能进行 k 次交易且不持股的最大利润
	return cash[k]
}

func main() {

}
