package main

func solveNQueens(n int) [][]string {
	// 初始化棋盘，全为 '.'
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	// 核心思路：
	// ● 逐行放置皇后，每行一个。
	// ● 使用三个布尔数组记录列、主对角线、副对角线的占用情况。
	// ● 递归函数 backtrack(row, n, board, cols, diag1, diag2, res) 尝试在第 row 行放置皇后。
	// 列占用
	cols := make([]bool, n)
	// 主对角线：row - col 为常数，范围 -(n-1) ~ n-1，偏移 n-1 作为索引
	diag1 := make([]bool, 2*n-1)
	// 副对角线：row + col 为常数，范围 0 ~ 2n-2
	diag2 := make([]bool, 2*n-1)
	res := [][]string{}

	backtrack(0, n, board, cols, diag1, diag2, &res)
	return res
}

// backtrack 递归放置皇后
// row: 当前要放置的行号（0-based）
// n: 棋盘大小
// board: 当前棋盘状态
// cols: 列占用标记
// diag1: 主对角线占用标记（索引 = row-col+n-1）
// diag2: 副对角线占用标记（索引 = row+col）
// res: 结果集指针
func backtrack(row, n int, board [][]byte, cols []bool, diag1, diag2 []bool, res *[][]string) {
	// 如果已经成功放置了 n 行，说明找到了一种解法
	if row == n {
		// 将当前棋盘 board 转换成字符串切片
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			solution[i] = string(board[i]) // 每一行 byte 转 string
		}
		*res = append(*res, solution) // 添加到结果集
		return
	}
	// 尝试当前行的每一列放置皇后
	for col := 0; col < n; col++ {
		// 计算当前格子对应的主对角线索引和副对角线索引
		d1 := row - col + n - 1 // 主对角线索引（偏移后）
		d2 := row + col         // 副对角线索引
		// 如果列、主对角线、副对角线任意一个已被占用，则跳过此列
		if cols[col] || diag1[d1] || diag2[d2] {
			continue
		}
		// 放置皇后
		board[row][col] = 'Q'
		cols[col] = true
		diag1[d1] = true
		diag2[d2] = true
		// 递归尝试下一行
		backtrack(row+1, n, board, cols, diag1, diag2, res)
		// 回溯，撤销当前皇后的放置
		board[row][col] = '.'
		cols[col] = false
		diag1[d1] = false
		diag2[d2] = false
	}
}

func main() {

}
