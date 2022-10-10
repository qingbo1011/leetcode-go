package main

import "sort"

func main() {

}

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
	//return quicksort(nums)
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := 0, len(nums)-1
	pivot := 0
	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]

	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}
