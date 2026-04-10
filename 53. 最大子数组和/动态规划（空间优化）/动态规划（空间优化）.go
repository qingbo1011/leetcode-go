package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	curMax := nums[0]
	for i := 1; i < len(nums); i++ {
		// if curMax >= 0 {
		// 	curMax = curMax + nums[i]
		// }else {
		// 	curMax = nums[i]
		// }
		// 写成if判断条件更容易理解：如果前面的子数组和是负数，那么最大子数组和当然没必要加上前面的子数组，直接以当前nums[i]开始算子数组即可
		// 优化后的写法就是
		curMax = max(curMax+nums[i], nums[i])
		if curMax > ans {
			ans = curMax
		}
	}

	return ans
}

func main() {

}
