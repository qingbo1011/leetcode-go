package main

func partition(s string) [][]string {
	n := len(s)
	// 1. 预处理 dp[i][j] 表示子串 s[i..j] 是否为回文
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 单字符回文
	}
	// 长度从 2 到 n
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	res := [][]string{}
	path := []string{}
	// 2. 回溯搜索所有分割方案
	backtrack(s, 0, &res, &path, dp)
	return res
}

// backtrack 回溯函数
// s: 原字符串
// start: 当前待分割的起始索引
// res: 结果集指针
// path: 当前路径指针
// dp: 预处理回文表
func backtrack(s string, start int, res *[][]string, path *[]string, dp [][]bool) {
	if start == len(s) {
		// 到达末尾，将当前路径拷贝一份加入结果集
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	// 尝试所有可能的结束位置
	for end := start; end < len(s); end++ {
		if dp[start][end] {
			// 当前子串是回文，加入路径
			*path = append(*path, s[start:end+1])
			// 递归处理剩余部分
			backtrack(s, end+1, res, path, dp)
			// 回溯：移除最后加入的子串
			*path = (*path)[:len(*path)-1]
		}
	}
}

func main() {

}
