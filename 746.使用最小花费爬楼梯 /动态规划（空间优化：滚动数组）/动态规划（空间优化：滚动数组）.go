package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	a, b := 0, 0 // dp[i-2] 和 dp[i-1]
	for i := 2; i <= len(cost); i++ {
		tmp := min(a+cost[i-2], b+cost[i-1])
		a, b = b, tmp
	}

	return b
}

func main() {

}
