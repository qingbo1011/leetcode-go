package main

func maxProfit(prices []int, fee int) int {
	cash := 0
	hold := -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		newCash := max(cash, hold+prices[i])
		newHold := max(hold, cash-prices[i]-fee)
		cash, hold = newCash, newHold
	}
	return cash // 最终不持股利润不低于持股
}

func main() {

}
