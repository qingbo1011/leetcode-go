package main

func rob(nums []int) int {
	// 状态定义：dp[i]表示前i间房屋（下标 0 到 i-1）能偷窃到的最高金额
	dp := make([]int, len(nums)+1)
	// 初始化：
	// dp[0] = 0（没有房屋）
	// dp[1] = nums[0]（只有第一间）
	dp[1] = nums[0]
	// 状态转移方程：dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	// 不偷第 i 间（下标 i-1）：dp[i-1]
	// 偷第 i 间：dp[i-2] + nums[i-1]
	// 取最大值：dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])，其中 i ≥ 2
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

func main() {

}
