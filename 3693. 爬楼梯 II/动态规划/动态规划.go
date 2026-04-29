package main

func climbStairs(n int, costs []int) int {
	// dp[i]表示从第0级台阶到达第i级台阶所需的最小总成本，其中 i ∈ [0, n]。
	dp := make([]int, n+1)
	// 初始条件：dp[0] = 0（起点，无需成本）

	// 状态转移方程：对于 i >= 1，可以从 i-1、i-2、i-3 跳过来：
	// dp[i] = min(
	//    dp[i-1] + costs[i-1] + 1²,      // 从 i-1 跳一步
	//    dp[i-2] + costs[i-1] + 2²,      // 从 i-2 跳两步
	//    dp[i-3] + costs[i-1] + 3²       // 从 i-3 跳三步
	// )
	// 注意：costs[i-1] 是目标台阶 i 的成本（因为 costs 下标从 0 对应第 1 级台阶）
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + costs[i-1] + 1*1
		if i-2 >= 0 {
			dp[i] = min(dp[i], dp[i-2]+costs[i-1]+2*2)
		}
		if i-3 >= 0 {
			dp[i] = min(dp[i], dp[i-3]+costs[i-1]+3*3)
		}
	}

	return dp[n]
}

func main() {

}
