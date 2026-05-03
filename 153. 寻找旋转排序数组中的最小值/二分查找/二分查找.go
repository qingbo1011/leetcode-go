package main

func findMin(nums []int) int {
	// 步骤：
	// 1. 初始化 left = 0，right = len(nums)-1。
	// 2. 当 left < right 时：
	//  ○ 计算 mid = left + (right-left)/2。
	//  ○ 如果 nums[mid] > nums[right]，最小值在右半部分，left = mid+1。
	//  ○ 否则，如果 nums[mid] < nums[right]，最小值在左半部分（包括 mid），right = mid。
	//  ○ 否则（nums[mid] == nums[right]），无法判断，right--。
	// 3. 返回 nums[left]。
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 如果中间值大于右边界，最小值在右半部分
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// 中间值小于右边界，最小值在左半部分（包括 mid）
			right = mid
		} else if nums[mid] == nums[right] {
			// 相等时无法判断，缩小右边界
			right--
		}
	}
	return nums[left]
}

func main() {

}
