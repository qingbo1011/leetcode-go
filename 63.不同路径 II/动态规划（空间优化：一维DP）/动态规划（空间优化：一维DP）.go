package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 起点或终点有障碍，直接返回 0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([]int, n)
	// 初始化第一行
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break // 遇到障碍，后面全为0（dp默认0）
		}
		dp[j] = 1
	}

	// 从第二行开始逐行更新
	for i := 1; i < m; i++ {
		// 每行第一个元素（第0列）单独处理：只能从上一行向下走
		if obstacleGrid[i][0] == 0 && dp[0] == 1 {
			dp[0] = 1
		} else {
			dp[0] = 0
		}
		// 从左到右更新当前行
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				// 新dp[j] = 上一行的dp[j]（即旧值） + 当前行左边的dp[j-1]（已更新）
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}

func main() {

}
