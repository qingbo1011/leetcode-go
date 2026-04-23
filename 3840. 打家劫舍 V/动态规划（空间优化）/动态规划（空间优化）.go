package main

func rob(nums []int, colors []int) int64 {
	// 状态定义：滚动变量
	// prev0: 考虑前 i-1 间房屋，且不偷第 i-1 间时的最大金额
	// prev1: 考虑前 i-1 间房屋，且偷第 i-1 间时的最大金额
	// 初始化：前 1 间房屋（下标 0）
	prev0 := int64(0)       // 不偷
	prev1 := int64(nums[0]) // 偷

	// 从第 2 间房屋开始（即前 i 间，i从2到n）
	for i := 2; i <= len(nums); i++ {
		// 不偷当前房屋（第 i 间，下标 i-1）
		cur0 := max(prev0, prev1)
		// 偷当前房屋
		var cur1 int64
		if colors[i-1] == colors[i-2] {
			// 颜色相同，前一间不能偷
			cur1 = int64(nums[i-1]) + prev0
		} else {
			// 颜色不同，前一间可偷可不偷
			cur1 = int64(nums[i-1]) + max(prev0, prev1)
		}
		prev0, prev1 = cur0, cur1
	}
	return max(prev0, prev1)
}

func main() {

}
