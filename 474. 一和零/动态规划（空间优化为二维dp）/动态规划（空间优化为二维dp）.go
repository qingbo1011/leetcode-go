package main

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	// dp[j][k]：表示在当前已经处理过的字符串中，使用不超过j个0和不超过k个1时，能得到的最大子集长度。
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始：dp[0][0] = 0，其余为 0（因为什么也不选时长度为 0）

	for _, str := range strs {
		zeros := strings.Count(str, "0")
		ones := len(str) - zeros
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				dp[j][k] = dp[j][k]
				if j >= zeros && k >= ones {
					dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1) // 这里+1，是因为选择“选这个字符串”，当然结果+1
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
