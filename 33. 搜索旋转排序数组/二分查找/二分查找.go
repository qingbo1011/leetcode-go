package main

func main() {

}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 开始二分
		middle := left + (right-left)>>1
		if nums[middle] == target {
			return middle
		}
		// 判断middle在左半部分还是右半部分
		if nums[middle] >= nums[left] { // middle在左半部分
			if nums[middle] < target { // target只可能在左半部分，和middle在同一边，可以使用二分
				left = middle + 1
			} else { // target可能在左半部分也可能在右半部分
				if target >= nums[left] { // target在左半部分，和middle在同一边，可以使用二分
					right = middle - 1
				} else { // target在右半部分，和middle不在同一边，不能使用二分，只能更新left
					left = middle + 1
				}
			}
		} else { // middle在右半部分
			if nums[middle] < target { // target可能在左半部分也可能在有半部分
				if target >= nums[left] { // target在左半部分,和middle不在同一边，不能使用二分只能更新right
					right = middle - 1
				} else { // target在右半部分，和middle在同一边，可以使用二分
					left = middle + 1
				}
			} else { // target只可能在右半部分，和middle在同一边，可以使用二分
				right = middle - 1
			}
		}
	}
	return -1
}
