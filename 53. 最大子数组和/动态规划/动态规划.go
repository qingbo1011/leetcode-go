package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	// dp数组含义：dp[i]表示以nums[i]结尾的连续子数组的最大和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// 状态转移方程：dp[i] = max(nums[i], dp[i-1] + nums[i])
	for i := 1; i < len(nums); i++ {
		// if dp[i-1] >= 0 {
		// 	dp[i] = dp[i-1] + nums[i]
		// }else {
		// 	dp[i] = nums[i]
		// }
		// 写成if判断条件更容易理解：如果前面的子数组和是负数，那么最大子数组和当然没必要加上前面的子数组，直接以当前nums[i]开始算子数组即可
		// 优化后的写法就是
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func main() {

}
