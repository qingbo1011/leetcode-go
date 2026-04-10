package main

func integerBreak(n int) int {
	// 定义dp[i]表示将正整数i拆分成至少两个正整数的和的最大乘积
	dp := make([]int, n+1)
	// 初始化dp
	dp[1] = 1
	// 转移方程：dp[i] = max(dp[i], j*(i-j), j*dp[i-j])  i≥3；1 ≤ j < i
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*dp[i-j], j*(i-j))
			// dp[i] = max(dp[i], dp[j]*(i-j), j*(i-j))	// 乘法的交换律：拆分j而保留i-j不拆分
		}
	}

	return dp[n]
}

func main() {

}
