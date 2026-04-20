package main

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	// 状态定义：dp[i][j][k]表示考虑前i个字符串（下标 0 ~ i-1），在最多使用j个0和k个1的条件下，能够得到的最大子集长度
	// 即最多能选多少个字符串。其中 0 ≤ i ≤ len(strs)，0 ≤ j ≤ m，0 ≤ k ≤ n
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(strs)+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	// 初始化：当i=0时，表示没有考虑任何字符串，此时无论容量如何，最大子集长度都为0。所以 dp[0][j][k] = 0 对所有j,k 成立。

	// 状态转移方程：dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeros][k-ones] + 1)
	// (若 j>=zeros 且 k>=ones)
	// 对于第i个字符串strs[i-1]，设它包含zeros个0和ones个1（ones = len(strs[i-1]) - zeros）。我们有两种选择：
	// 不选这个字符串：那么当前最大子集长度就等于前 i-1 个字符串在相同容量下的结果，即 dp[i-1][j][k]。
	// 选这个字符串（前提是 j >= zeros 且 k >= ones）：此时长度 = dp[i-1][j-zeros][k-ones] + 1。
	for i := 1; i <= len(strs); i++ {
		// 统计当前字符串的0和1个数(注意这里遍历的i表示前i个字符串，因此当前字符串是strs[i-1])
		zeros := strings.Count(strs[i-1], "0")
		ones := len(strs[i-1]) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= zeros && k >= ones {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-zeros][k-ones]+1) // 这里+1，是因为选择“选这个字符串”，当然结果+1
				}
			}
		}
	}

	// 最终答案：经过所有字符串的决策后，dp[len(strs)][m][n]即为所求
	return dp[len(strs)][m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
