package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1 // 初始化为-1
		}
	}
	// 定义四个边界：top（上边界）、bottom（下边界）、left（左边界）、right（右边界）。
	top, bottom, left, right := 0, m-1, 0, n-1
	for top <= bottom && left <= right {
		// 在每一圈中，按顺序遍历：
		// 1.从左到右遍历上边界（top行，列从left到right）
		// 2.从上到下遍历右边界（right列，行从top+1到bottom）
		// 3.如果top<bottom,从右到左遍历下边界（bottom行，列从right-1到left） && left<right
		// 4.如果left<right，从下到上遍历左边界（left列，行从bottom-1到top+1）
		for j := left; j <= right; j++ {
			if head != nil {
				matrix[top][j] = head.Val
				head = head.Next
			}
		}
		for i := top + 1; i <= bottom; i++ {
			if head != nil {
				matrix[i][right] = head.Val
				head = head.Next
			}
		}
		// 使用top < bottom && left < right确保只有多于一行一列时才遍历下边和左边，否则会重复
		if top < bottom {
			for j := right - 1; j >= left; j-- {
				if head != nil {
					matrix[bottom][j] = head.Val
					head = head.Next
				}
			}
		}
		if left < right {
			for i := bottom - 1; i >= top+1; i-- {
				if head != nil {
					matrix[i][left] = head.Val
					head = head.Next
				}
			}
		}
		// 每遍历完一圈，收缩边界：top++，bottom--，left++，right--
		top++
		bottom--
		left++
		right--
	}

	return matrix
}

func main() {

}
