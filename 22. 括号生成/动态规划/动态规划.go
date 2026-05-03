package main

func generateParenthesis(n int) []string {
	// 任何有效括号组合都可以表示为 (A)B 的形式，其中 A 和 B 都是有效的括号组合（可能为空）。
	if n == 0 {
		return []string{}
	}
	// dp[i]存储i对括号的所有有效组合
	dp := make([][]string, n+1)
	dp[0] = []string{""} // 0对括号，只有空字符串
	for i := 1; i <= n; i++ {
		dp[i] = []string{}
		// 枚举左括号内包含的括号对数 a，那么右半部分为 i-1-a 对
		for a := 0; a < i; a++ {
			b := i - 1 - a
			for _, left := range dp[a] {
				for _, right := range dp[b] {
					// 组合成 (A)B
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}

	return dp[n]
}

func main() {

}
