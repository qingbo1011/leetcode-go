package main

func totalNQueens(n int) int {
	// 列占用标记
	cols := make([]bool, n)
	// 主对角线占用标记（row - col + n - 1）
	diag1 := make([]bool, 2*n-1)
	// 副对角线占用标记（row + col）
	diag2 := make([]bool, 2*n-1)
	count := 0
	backtrack(0, n, cols, diag1, diag2, &count)
	return count
}

// backtrack 递归放置皇后，统计合法方案数
// row: 当前要放置的行号（0-based）
// n: 棋盘大小
// cols, diag1, diag2: 占用标记数组
// count: 计数器指针
func backtrack(row, n int, cols, diag1, diag2 []bool, count *int) {
	if row == n {
		// 成功放置完 n 行，找到一种方案
		*count++
		return
	}
	// 尝试当前行的每一列
	for col := 0; col < n; col++ {
		d1 := row - col + n - 1
		d2 := row + col
		if cols[col] || diag1[d1] || diag2[d2] {
			continue
		}
		// 放置皇后
		cols[col] = true
		diag1[d1] = true
		diag2[d2] = true
		// 递归下一行
		backtrack(row+1, n, cols, diag1, diag2, count)
		// 回溯撤销
		cols[col] = false
		diag1[d1] = false
		diag2[d2] = false
	}
}

func main() {

}
