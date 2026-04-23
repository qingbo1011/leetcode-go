package main

func minCapability(nums []int, k int) int {
	// 题目等价于：在不相邻的约束下，选至少 k 个数，使得这些数的最大值最小。
	// 解题思路：二分答案 + 贪心验证：
	// 单调性：设 f(X) 表示当窃取能力为 X 时，能否偷到至少 k 间房屋。
	// 如果 X 可行，那么任何 Y > X 也可行（因为限制放宽了）。
	// 因此答案具有单调性，可以二分查找最小的可行 X。
	// 贪心策略：从左到右扫描数组，只要当前房屋金额 ≤ cap，就偷它，然后跳过下一间（因为不能相邻），继续扫描。这样能最大化偷的数量。
	// 为什么贪心正确？因为如果当前房屋可选，不偷它而考虑后面的，最多也只能偷同样数量（因为跳过一个，后面不一定更优），且可能因为跳过而浪费机会。所以遇到可偷的立即偷是最优的。

	// 二分边界：
	// 左边界 left：min(nums)（因为至少偷一间，能力至少为最小值）
	// 右边界 right：max(nums)（能力最大不会超过最大值）
	// 答案一定在 nums 中的某个值，所以二分区间可设为 [min(nums), max(nums)]。
	low, high := nums[0], nums[0]
	for _, v := range nums {
		if v < low {
			low = v
		}
		if v > high {
			high = v
		}
	}

	for low < high {
		mid := (low + high) / 2
		if canPick(nums, k, mid) {
			high = mid // 尝试更小的能力
		} else {
			low = mid + 1
		}
	}
	return low
}

// 验证函数canPick(cap)：给定一个能力cap，判断是否能偷至少k间不相邻的房屋，且每间金额≤cap
func canPick(nums []int, k int, cap int) bool {
	cnt := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= cap {
			cnt++
			i += 2 // 偷了当前，下一间不能偷
		} else {
			i++
		}
	}
	return cnt >= k
}

func main() {

}
