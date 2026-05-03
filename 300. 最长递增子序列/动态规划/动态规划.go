package main

import "fmt"

func lengthOfLIS(nums []int) int {
	// 状态定义：dp[i]表示以第i个元素（nums[i]）结尾的最长严格递增子序列的长度
	dp := make([]int, len(nums))
	// 初始化：每个dp[i]初始为1，因为每个元素自身可以作为一个长度为1的递增子序列
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	// 状态转移方程：对于每个i，考虑所有j< i：
	//● 如果 nums[j]<nums[i]，那么可以将nums[i] 接在以nums[j]结尾的递增子序列后面，形成长度为dp[j] + 1的子序列
	//● dp[i]取所有可能中的最大值，并至少为 （自身单独构成子序列）
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 最终答案：max(dp[0], dp[1], ..., dp[n-1])，即所有dp[i]中的最大值
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
