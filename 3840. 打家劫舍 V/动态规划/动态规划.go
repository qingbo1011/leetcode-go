package main

func rob(nums []int, colors []int) int64 {
	// 状态定义：
	// dp[i][0]：考虑前 i 间房屋，且第 i 间（下标 i-1）不偷时，能获得的最大金额。
	// dp[i][1]：考虑前 i 间房屋，且第 i 间（下标 i-1）偷时，能获得的最大金额。
	dp := make([][]int64, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int64, 2)
	}
	// 初始化：dp[0][0] = 0，dp[0][1]无意义（或设为负无穷），因为前0间房屋不存在第0间。通常我们直接从i=1开始。

	// 状态转移：对于 i >= 1：
	// 不偷第 i 间：dp[i][0] = max(dp[i-1][0], dp[i-1][1])（前一间可偷可不偷）
	// 偷第 i 间：需要考虑第 i 间与第 i-1 间的颜色是否相同
	//  若 colors[i-1] == colors[i-2]（注意索引偏移），则前一间不能偷，dp[i][1] = nums[i-1] + dp[i-1][0]
	//  否则，dp[i][1] = nums[i-1] + max(dp[i-1][0], dp[i-1][1])
	for i := 1; i <= len(nums); i++ {
		// 计算不偷第i间能获得的最大金额
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		// 计算偷第i间能获得的最大金额
		if i >= 2 && colors[i-1] == colors[i-2] {
			// 颜色相同前一间不能偷
			dp[i][1] = dp[i-1][0] + int64(nums[i-1])
		} else {
			// 能偷前一间
			dp[i][1] = max(dp[i-1][0], dp[i-1][1]) + int64(nums[i-1])
		}
	}

	// 最终答案：max(dp[n][0], dp[n][1])
	return max(dp[len(nums)][0], dp[len(nums)][1])
}

func main() {

}
