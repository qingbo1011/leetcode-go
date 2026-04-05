package main

import "fmt"

func trap(height []int) int {
	water := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			// 左侧较矮，处理左侧
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water = water + leftMax - height[left]
			}
			left++
		} else {
			// 右侧较矮，处理右侧
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water = water + rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
