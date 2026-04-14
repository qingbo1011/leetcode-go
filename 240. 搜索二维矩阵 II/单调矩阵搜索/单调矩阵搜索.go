package main

func searchMatrix(matrix [][]int, target int) bool {
	// 由于该矩阵每行的元素从左到右升序排列；每列的元素从上到下升序排列
	// 也就是说向下变大，向左变小。我们设置一个起点在matrix[0][len(matrix[0])-1]
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			// 向下移动
			i++
		} else if matrix[i][j] > target {
			// 向左移动
			j--
		}
	}
	return false
}

func main() {

}
