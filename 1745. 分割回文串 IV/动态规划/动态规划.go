package main

func checkPartitioning(s string) bool {
	n := len(s)
	if n < 3 {
		return false
	}
	// 状态定义：isPal[i][j] 表示字符串 s 的子串 s[i..j]（闭区间）是否为回文串
	isPal := make([][]bool, n)
	// 初始化：
	// ● 对于所有 i，isPal[i][i] = true（单个字符是回文）
	// ● 对于长度 len = 2，isPal[i][i+1] = (s[i] == s[i+1])
	for i := 0; i < n; i++ {
		isPal[i] = make([]bool, n)
		isPal[i][i] = true // 单字符回文
	}

	// 状态转移方程：对于长度 len ≥ 3 的子串 s[i..j]（j = i+len-1）：
	// isPal[i][j] = (s[i] == s[j]) && isPal[i+1][j-1]。即两端字符相等且内部子串为回文
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || isPal[i+1][j-1]) {
				isPal[i][j] = true
			}
		}
	}
	// 最终答案：枚举两个分割点 i（第一个子串结束位置+1）和 j（第二个子串结束位置+1），其中 1 ≤ i ≤ n-2，i+1 ≤ j ≤ n-1。
	// 若存在 i, j 使得以下三个条件同时成立：
	// ● isPal[0][i-1]（第一个子串回文）
	// ● isPal[i][j-1]（第二个子串回文）
	// ● isPal[j][n-1]（第三个子串回文）
	for i := 1; i <= n-2; i++ {
		if !isPal[0][i-1] {
			continue // 第一个子串不是回文，跳过
		}
		// 第二个分割点 j：第二个子串为 s[i..j-1]，j 范围 i+1 到 n-1
		for j := i + 1; j <= n-1; j++ {
			if isPal[i][j-1] && isPal[j][n-1] {
				return true
			}
		}
	}
	return false
}

func main() {

}
