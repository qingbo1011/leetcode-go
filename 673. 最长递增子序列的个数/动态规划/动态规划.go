package main

func findNumberOfLIS(nums []int) int {
	// 状态定义：
	//● dp[i]：以 nums[i] 结尾的最长递增子序列的长度
	//● cnt[i]：以 nums[i] 结尾的长度为dp[i]的递增子序列个数
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	// 初始化：每个位置dp[i] = 1，cnt[i] = 1（自身单独一个子序列）
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
	}
	// 状态转移方程：遍历所有 i，对于每个 j < i：
	//● 如果 nums[j] < nums[i]：
	//  ○ 若 dp[j] + 1 > dp[i]，则更新 dp[i] = dp[j] + 1，cnt[i] = cnt[j]（新长度覆盖旧长度）
	//  ○ 若 dp[j] + 1 == dp[i]，则 cnt[i] += cnt[j]（累加相同长度的不同来源）
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] = cnt[i] + cnt[j]
				}
			}
		}
	}
	// 最终答案：找出最大的dp[i]（记为 maxLen），累加所有dp[i]== maxLen的cnt[i]，即为所求。
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		maxLength = max(maxLength, dp[i])
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLength {
			count = count + cnt[i]
		}
	}
	return count
}

func main() {

}
