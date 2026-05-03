package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	// 长度剪枝
	if len(word) > m*n {
		return false
	}

	// 频率剪枝
	boardCount := make([]int, 128) // ASCII 码范围
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			boardCount[board[i][j]]++
		}
	}
	wordCount := make([]int, 128)
	for i := 0; i < len(word); i++ {
		wordCount[word[i]]++
		if wordCount[word[i]] > boardCount[word[i]] {
			return false // 某字符需求超过网格可用量
		}
	}

	// 端点优化：选择出现次数较少的端作为开始
	// 如果首字符出现次数 > 尾字符，则反转单词，从尾部搜索能减少分支
	if boardCount[word[0]] > boardCount[word[len(word)-1]] {
		word = reverse(word)
	}

	// 深度优先搜索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// reverse 反转字符串
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// dfs 深度优先搜索，从 (i,j) 开始匹配 word[idx:]
func dfs(board [][]byte, i, j int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}
	// 边界检查或字符不匹配
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[idx] {
		return false
	}
	// 标记当前格子为已访问（临时修改）
	temp := board[i][j]
	board[i][j] = '#'

	// 四个方向尝试（方向顺序可调整，但无所谓）
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
