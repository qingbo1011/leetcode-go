package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 起点或终点有障碍，直接返回 0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// 初始化二维dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		// 一旦遇到障碍，后面的都为0（默认已是 0）
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		// 一旦遇到障碍，后面的都为0（默认已是 0）
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	// 状态转移只判断obstacleGrid[i][j]，而不判断上方和左方格子.因为障碍物的影响已经通过dp数组的值间接体现了：
	// 如果上方格子(i-1, j)有障碍，那么dp[i-1][j]在初始化或之前的转移中一定为 0（因为障碍物格子不会被赋值为路径数，且后续也无法到达）。
	// 同理，左方格子(i, j-1)如果有障碍，dp[i][j-1]也为 0。
	// 因此，dp[i][j] = dp[i-1][j] + dp[i][j-1] 自动加上了0，不会错误地累加路径。
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {

}
