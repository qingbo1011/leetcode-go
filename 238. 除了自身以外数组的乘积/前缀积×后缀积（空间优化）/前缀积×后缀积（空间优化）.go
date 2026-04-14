package main

import "fmt"

func productExceptSelf(nums []int) []int {
	// 优化步骤：
	//1. 去掉 left 数组：直接用 ans 数组存储左边乘积。
	//2. 去掉 right 数组：用一个变量 right 从右向左累乘，同时更新 ans[i] *= right。
	right := 1
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right = right * nums[i]
	}

	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
