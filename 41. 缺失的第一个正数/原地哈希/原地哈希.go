package main

import "fmt"

func firstMissingPositive(nums []int) int {
	// 将nums[i]放到正确的位置nums[i]-1处
	for i := 0; i < len(nums); i++ {
		for (nums[i] >= 1 && nums[i] <= len(nums)) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// 寻找第一个不匹配的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
