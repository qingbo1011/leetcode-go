package main

import "sort"

func main() {

}
func searchRange(nums []int, target int) []int {
	leftBorder := sort.SearchInts(nums, target)
	if leftBorder == len(nums) || nums[leftBorder] != target { // 说明target不存在数组中
		return []int{-1, -1}
	}
	// target存在数组中,开始找rightBorder
	rightBorder := sort.SearchInts(nums, target+1) - 1
	return []int{leftBorder, rightBorder}
}
