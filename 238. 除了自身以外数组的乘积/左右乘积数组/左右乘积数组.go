package main

import "fmt"

func productExceptSelf(nums []int) []int {
	// left[i]表示nums[i]左边的乘积
	// right[i]表示nums[i]右边的乘积
	// ans[i] = left[i]*right[i]
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	ans := make([]int, len(nums))

	left[0], right[len(nums)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
