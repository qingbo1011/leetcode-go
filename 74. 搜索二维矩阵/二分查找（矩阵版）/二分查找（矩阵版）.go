package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 将二维矩阵展开成一维数组（逻辑上），索引idx对应矩阵位置 (idx/n, idx%n)
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else if val > target {
			right = mid - 1
		}
	}

	return false
}

func main() {

}
