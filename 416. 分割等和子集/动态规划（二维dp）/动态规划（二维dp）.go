package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	// 总和为奇数则一定不能割成两个子集，使得两个子集的元素和相等
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// 只要能在nums中任取i个数字使得其和为target，则说明可以割成两个子集，使得两个子集的元素和相等

	// dp数组含义：dp[i][j] -> 考虑前i个数字（即下标0到i-1），能否选出若干个使其总和为j
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]bool, target+1)
	}
	// 初始化:
	// dp[0][0] = true：前0个数字可以凑出和为 0（什么都不选）
	// dp[0][j] = false（j > 0）：前0个数字无法凑出正数和
	dp[0][0] = true
	// 状态转移方程：dp[i][j] = dp[i-1][j] || (j >= nums[i-1] && dp[i-1][j - nums[i-1]])
	// 对于第i个数字（值为 nums[i-1]），有两种选择：
	// 不选：能否凑出j完全取决于前i-1个数字，即 dp[i][j] = dp[i-1][j]
	// 选：需要满足j>=nums[i-1]，并且前i-1个数字能凑出j-nums[i-1]，即dp[i-1][j-nums[i-1]]为真。此时dp[i][j]也为真
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j] || (j >= nums[i-1] && dp[i-1][j-nums[i-1]])
		}
	}

	return dp[len(nums)][target]
}

func main() {

}
