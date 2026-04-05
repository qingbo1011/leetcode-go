package main

import "fmt"

func maxArea(height []int) int {
	result := 0
	// 左右指针从两端开始向内移动，直至重合
	left, right := 0, len(height)-1
	for left < right {
		water := (right - left) * min(height[left], height[right])
		if water > result {
			result = water
		}
		// 贪心思想：由于是两端向内移动，宽度会变短，因此需要移动较矮的一方，这样才能保证可能接的水更多
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
