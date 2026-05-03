package main

func search(nums []int, target int) int {
	// ● 每次取中间位置 mid，比较 nums[mid] 与 nums[left]。
	// ● 如果 nums[left] <= nums[mid]，说明左半部分 [left, mid] 是有序的（可能包含最大值和最小值交界，但该条件保证左边有序）。
	//  ○ 若 target 在 [nums[left], nums[mid]] 范围内，则缩小到左半部分 [left, mid-1]。
	//  ○ 否则，搜索右半部分 [mid+1, right]。
	// ● 否则，右半部分 [mid, right] 是有序的。
	//  ○ 若 target 在 [nums[mid], nums[right]] 范围内，则缩小到右半部分 [mid+1, right]。
	//  ○ 否则，搜索左半部分 [left, mid-1]。
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 左半部分有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1 // target 在左半部分
			} else {
				left = mid + 1 // target 在右半部分
			}
		} else {
			// 右半部分有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1 // target 在右半部分
			} else {
				right = mid - 1 // target 在左半部分
			}
		}
	}

	return -1
}

func main() {

}
