package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	// 递归公式：dp[i] = dp[i-1] + dp[i-2]（i>=3）
	// 由于一次只能走1步或2步，所以走i阶的走法有dp[i-1] + dp[i-2]（走完dp[i-1]再走一步，走完dp[i-2]再走两步）
	// dp[i-1]已经包含了所有“最后一步是1阶”的情况，dp[i-2]包含了所有“最后一步是2阶”的情况，两者互斥且完备。因此递推公式就是 dp[i]=dp[i-1]+dp[i-2]
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {

}
