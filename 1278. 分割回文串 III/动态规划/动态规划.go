package main

import "math"

func palindromePartition(s string, k int) int {
	n := len(s)
	// 状态定义： cost[l][r] 表示将子串 s[l..r]（闭区间）变成回文串所需的最小修改次数。
	// 初始化：cost[l][r]：当 l >= r 时，cost[l][r] = 0；否则通过递推计算。

	// 1. 预处理 cost[l][r] 表示子串 s[l..r] 变成回文的最小修改次数
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		// cost[i][i] 默认为 0
	}
	// 状态转移方程：预处理所有子串的修改代价：
	// 对于长度 len 从 2 到 n，l 从 0 到 n-len，r = l+len-1：
	// cost[l][r] = cost[l+1][r-1] + (s[l] == s[r] ? 0 : 1)
	// 注意：当 len 为奇数时，中心字符无需配对，不影响公式。
	// 按子串长度递增顺序计算，保证 cost[l+1][r-1] 已计算
	for length := 2; length <= n; length++ {
		for l := 0; l+length-1 < n; l++ {
			r := l + length - 1
			if s[l] == s[r] {
				cost[l][r] = cost[l+1][r-1] // 两端相等，无需额外修改
			} else {
				cost[l][r] = cost[l+1][r-1] + 1 // 两端不等，需修改其中一端
			}
		}
	}

	// 状态定义：dp[i][j] 表示将字符串 s 的前 i 个字符（即 s[0:i]）分割成 j 个非空回文子串所需的最小修改次数。
	dp := make([][]int, n+1)
	// 初始化：dp[0][0] = 0，表示空串分成 0 段无需修改。
	// 对于 i > 0 或 j > 0，dp[i][j] 初始化为一个足够大的数
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0 // 空串分成0段，修改0次

	// 状态转移方程： 主DP转移：
	// 对于每个 i（1 ≤ i ≤ n）和每个 j（1 ≤ j ≤ k 且 j ≤ i）：
	// dp[i][j] = min_{t = j-1}^{i-1} ( dp[t][j-1] + cost[t][i-1] )
	// 其中 t 表示前 t 个字符已被分成 j-1 段，最后一段为 s[t..i-1]。
	for i := 1; i <= n; i++ {
		// 最多分成 min(i, k) 段
		maxJ := min(i, k)
		for j := 1; j <= maxJ; j++ {
			// 枚举最后一段的起始索引 t（前 t 个字符已经分成 j-1 段）
			// t 至少为 j-1，因为前 j-1 段至少需要 j-1 个字符
			for t := j - 1; t < i; t++ {
				if dp[t][j-1] != math.MaxInt {
					// 最后一段是 s[t..i-1]，其修改代价为 cost[t][i-1]
					cur := dp[t][j-1] + cost[t][i-1]
					if cur < dp[i][j] {
						dp[i][j] = min(dp[i][j], dp[t][j-1]+cost[t][i-1])
					}
				}
			}
		}
	}
	return dp[n][k]
}

func main() {

}
