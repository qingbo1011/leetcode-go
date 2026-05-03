package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	// 找左边界：第一个等于 target 的位置
	left := findLeftBound(nums, target)
	// 如果左边界无效或值不等于 target，直接返回 [-1,-1]
	if left == -1 || nums[left] != target {
		return []int{-1, -1}
	}
	// 找右边界：最后一个等于 target 的位置
	right := findRightBound(nums, target)
	return []int{left, right}
}

// findLeftBound 返回第一个 >= target 的索引，若不存在则返回 -1
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	idx := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			idx = mid       // 记录候选
			right = mid - 1 // 继续向左找更左的
		} else {
			left = mid + 1
		}
	}
	return idx
}

// findRightBound 返回最后一个 <= target 的索引，若不存在则返回 -1
func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	idx := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			idx = mid      // 记录候选
			left = mid + 1 // 继续向右找更右的
		} else {
			right = mid - 1
		}
	}
	return idx
}

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Ints(s)
	fmt.Println(s) // [-3 1 3 3 5 6 8 8 9 15 45 66]
	fmt.Println(sort.SearchInts(s, 5))
}
