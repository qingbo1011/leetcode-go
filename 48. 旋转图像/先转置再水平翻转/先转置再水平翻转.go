package main

func rotate(matrix [][]int) {
	n := len(matrix)
	// 先转置再水平翻转的方法仅适用n*n的矩阵，如果是m*n的矩阵则只能使用额外空间，再去一一对应。
	// 1.将矩阵转置（matrix[i][j]与matrix[j][i]交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 2.对每一行进行反转（左右镜像）
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func main() {

}
