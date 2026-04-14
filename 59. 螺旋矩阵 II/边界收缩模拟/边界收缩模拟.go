package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 定义四个边界：top（上边界）、bottom（下边界）、left（左边界）、right（右边界）。
	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1
	for top <= bottom && left <= right {
		// 在每一圈中，按顺序遍历：
		// 1.从左到右遍历上边界（top行，列从left到right）
		// 2.从上到下遍历右边界（right列，行从top+1到bottom）
		// 3.如果top<bottom,从右到左遍历下边界（bottom行，列从right-1到left） && left<right
		// 4.如果left<right，从下到上遍历左边界（left列，行从bottom-1到top+1）
		for j := left; j <= right; j++ {
			matrix[top][j] = num
			num++ // 每次写入后num++
		}
		for i := top + 1; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		// 使用top < bottom && left < right确保只有多于一行一列时才遍历下边和左边，否则会重复
		if top < bottom {
			for j := right - 1; j >= left; j-- {
				matrix[bottom][j] = num
				num++
			}
		}
		if left < right {
			for i := bottom - 1; i >= top+1; i-- {
				matrix[i][left] = num
				num++
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
