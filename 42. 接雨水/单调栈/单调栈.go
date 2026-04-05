package main

import "fmt"

func trap(height []int) int {
	ans := 0
	// 单调栈(栈底大)。注意栈中存储索引，保持栈内索引对应的高度递减
	stack := make([]int, 0, len(height))
	for i, h := range height {
		// 出现凹槽，即高-低-高结构
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			bottomHeight := height[stack[len(stack)-1]] // 记录出栈元素高度（即高-低-高结构中低柱子的高度）
			// 弹出元素
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 左边没有柱子，无法形成凹槽
				break
			}
			leftIdx := stack[len(stack)-1] // 左边柱子索引
			// 宽度 = 右边柱子索引 - 左边柱子索引 - 1
			width := i - leftIdx - 1
			// 高度 = min(左边柱子高度, 右边柱子高度) - 底部高度
			heightDiff := min(height[leftIdx], h) - bottomHeight
			ans = ans + width*heightDiff
		}
		// 当前索引入栈
		stack = append(stack, i)
	}

	return ans
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
