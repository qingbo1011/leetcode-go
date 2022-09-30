package main

func main() {

}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target { // 更新left
			left = middle + 1
		} else { // 更新right
			right = middle - 1
		}
	}
	return left
}
