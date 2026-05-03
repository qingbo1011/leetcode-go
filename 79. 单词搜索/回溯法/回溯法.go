package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	// 遍历网格的每个单元格作为起点，尝试匹配单词的第一个字母
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// dfs 深度优先搜索，判断从 (i,j) 出发是否能匹配 word 从 idx 开始的后缀
func dfs(board [][]byte, i, j int, word string, idx int) bool {
	// 如果已经匹配完所有字符，成功
	if idx == len(word) {
		return true
	}
	// 边界检查或字符不匹配
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[idx] {
		return false
	}
	// ● 从当前单元格出发，向四个方向（上、下、左、右）递归探索，每次匹配单词的下一个字母。
	// ● 使用一个临时标记（如将 board 中的字符替换为 '#'）表示该单元格已访问，递归结束后恢复原字符，实现回溯。
	// 保存当前字符并标记为已访问（用一个临时字符覆盖）
	temp := board[i][j]
	board[i][j] = '#'
	// 四个方向递归尝试
	if dfs(board, i+1, j, word, idx+1) ||
		dfs(board, i-1, j, word, idx+1) ||
		dfs(board, i, j+1, word, idx+1) ||
		dfs(board, i, j-1, word, idx+1) {
		return true
	}
	// 回溯：恢复原字符
	board[i][j] = temp
	return false
}

func main() {

}
