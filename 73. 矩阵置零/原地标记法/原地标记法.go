package main

func setZeroes(matrix [][]int) {
	firstRowZero, firstColZero := false, false
	// 先遍历第一行和第一列，判断它们是否包含0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	// 遍历除第一行、第一列外的所有元素，如果matrix[i][j]==0，则将 matrix[i][0]和matrix[0][j]置为0（作为标记）
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 再次遍历除第一行、第一列外的所有元素，如果matrix[i][0]==0或matrix[0][j]==0，则将matrix[i][j]置为0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			// 当前位置[i][j]的最左方或最上方为0
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 根据第一步的标记，如果firstRowZero为真，则将第一行全部置为0；如果firstColZero为真，则将第一列全部置为0
	if firstRowZero {
		// 第一行全部置为0
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		// 第一列全部置为0
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {

}
