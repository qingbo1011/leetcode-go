package main

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	// 初始化dp(第0行全是1)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ { // 从第1行开始更新
		for j := 1; j < n; j++ { // 每行从第1列开始
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

func main() {

}
