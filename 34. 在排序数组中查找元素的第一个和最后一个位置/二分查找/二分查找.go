package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Ints(s)
	fmt.Println(s) // [-3 1 3 3 5 6 8 8 9 15 45 66]
	fmt.Println(sort.SearchInts(s, 5))
}

func searchRange(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	} else if rightBorder-leftBorder >= 2 {
		return []int{leftBorder + 1, rightBorder - 1}
	} else {
		return []int{-1, -1}
	}
}

/*
  二分查找，寻找target的左边界leftBorder（不包括target）
  考虑到leftBorder没有被赋值的情况（即target在数组范围的右边，例如数组[3,3],target为4），用来处理情况一
*/
func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	leftBorder := -2 // leftBorder为-2，表示leftBorder没有被赋值（情况一）
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] < target { // target在有区间，更新区间为[middle+1,right]
			left = middle + 1
		} else {
			right = middle - 1
			leftBorder = right
		}
	}
	return leftBorder
}

/*
  二分查找，寻找target的右边界（不包括target）
  考虑到rightBorder没有被赋值的情况（即target在数组范围的左边，例如数组[3,3]，target为2），用来处理情况一
*/
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightBorder := -2 // rightBorder为-2，表示rightBorder没有被赋值（情况一）
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target { // 更新right（target在左区间，所以更新区间为[left,middle-1]）
			right = middle - 1
		} else if nums[middle] == target { // 找到了，边界为middle+1（但是此时不能break！要继续循环）
			rightBorder = middle + 1
			left = middle + 1
		} else if nums[middle] < target { // 更新left（target在右区间，所以更新区间为[middle+1,right]）
			left = middle + 1
		}
		//if nums[middle] > target {
		//	right = middle - 1 // target 在左区间，所以[left, middle - 1]
		//} else { // 当nums[middle] == target的时候，更新left，这样才能得到target的右边界
		//	left = middle + 1
		//	rightBorder = left
		//}
	}
	return rightBorder
}
