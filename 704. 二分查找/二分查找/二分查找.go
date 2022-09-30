package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
