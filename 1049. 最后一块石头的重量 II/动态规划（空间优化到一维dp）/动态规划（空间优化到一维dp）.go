package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, w := range stones {
		sum += w
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, w := range stones {
		for j := target; j >= w; j-- {
			if dp[j-w] {
				dp[j] = true
			}
		}
	}

	// 找到最大的可达重量
	maxWeight := 0
	for j := target; j >= 0; j-- {
		if dp[j] {
			maxWeight = j
			break
		}
	}
	return sum - 2*maxWeight
}

func main() {

}
