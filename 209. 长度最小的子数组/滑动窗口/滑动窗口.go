package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	res := math.MaxInt
	left, i := 0, 0
	sum := 0
	for i < length {
		sum = sum + nums[i]
		for sum >= target {
			res = min(res, i-left+1)
			sum = sum - nums[left]
			left++
		}
		i++
	}
	//for i, num := range nums {
	//	sum = sum + num
	//	for sum >= target {
	//		res = min(res, i-left+1)
	//		sum = sum - nums[left]
	//		left++
	//	}
	//}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
