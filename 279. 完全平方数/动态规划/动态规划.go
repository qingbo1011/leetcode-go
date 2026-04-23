package main

func numSquares(n int) int {
	// 状态定义：dp[i] 表示组成整数i所需的最少完全平方数个数。
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = n + 1 // 初始化为较大值(方便状态转移时通过min赋值)
	}
	// 初始化：dp[0]=0（0 不需要任何完全平方数）
	dp[0] = 0
	// 状态转移方程：对于每个i，尝试减去一个完全平方数j*j（j*j ≤ i），则：
	// dp[i] = min(dp[i], dp[i - j*j] + 1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func main() {

}
