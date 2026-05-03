package main

func minCut(s string) int {
	n := len(s)
	// 提前使用动态规划预处理所有子串是否为回文，得到二维布尔数组isPal。方便后续判断s[j:i]是否为回文
	isPal := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPal[i] = make([]bool, n)
		isPal[i][i] = true // 单字符回文
	}
	// 长度从 2 开始
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || isPal[i+1][j-1]) {
				isPal[i][j] = true
			}
		}
	}

	// 这里开始用动态规划求最小分割次数
	// 状态定义：dp[i]表示字符串s的前i个字符（即 s[0:i]）分割成若干个回文子串所需的最少分割次数。约定 dp[0] = -1（空串无需分割，便于转移）。
	dp := make([]int, n+1)
	// 初始化：
	// ● dp[0] = -1
	// ● 对于 i >= 1，dp[i] 初始化为 i（最坏情况每个字符单独分割，需 i 次分割），但实际会在转移中更新。
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	// 状态转移方程：对于每个位置 i（1 ≤ i ≤ n），枚举所有可能的分割点 j（0 ≤ j < i）：
	// ● 如果子串 s[j:i] 是回文串，则 dp[i] = min(dp[i], dp[j] + 1)。
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if isPal[j][i-1] { // 这里是dp[i]的i，代表前i个字符。因此对于到isPal的s[j:i]的索引i，需要i-1
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n]
}

func main() {

}
