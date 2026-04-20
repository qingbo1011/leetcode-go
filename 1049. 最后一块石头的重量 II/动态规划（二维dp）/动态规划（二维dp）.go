package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	// 将石头分成两堆，使两堆重量之差最小。等价于：选出一组石头，使其总重量尽可能接近sum/2，则剩余重量=sum-2*选出的重量
	// 总重量 sum，我们选出的一组石头总重量为 maxWeight（不超过 sum/2 的最大可能值）。
	// 那么剩下的另一组石头总重量为sum - maxWeight。
	// 粉碎过程中，两组石头相互抵消，最后剩下的重量为两堆重量之差的绝对值：
	// |(sum - maxWeight) - maxWeight| = |sum - 2*maxWeight|。
	// 由于 maxWeight ≤ sum/2，所以 sum - 2*maxWeight ≥ 0，绝对值可直接去掉，得到sum - 2*maxWeight
	sum := 0
	for _, stone := range stones {
		sum = sum + stone
	}
	target := sum / 2
	// dp数组含义：dp[i][j] -> 考虑前i个数字（即下标0到i-1），能否选出若干个使其总和为j
	dp := make([][]bool, len(stones)+1)
	// 初始化：
	// dp[0][0] = true：前 0 块石头可以凑出重量 0。
	// 其他 dp[0][j] = false。
	for i := 0; i < len(stones)+1; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	// 状态转移方程：对于第 i 块石头（重量 w = stones[i-1]）：dp[i][j] = dp[i-1][j] || (j >= w && dp[i-1][j-w])
	// 不选这块石头：dp[i][j] = dp[i][j] || dp[i-1][j]
	// 选这块石头（需 j ≥ w）：dp[i][j] = dp[i][j] || dp[i-1][j-w]
	for i := 1; i <= len(stones); i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j] || (j-stones[i-1] >= 0 && dp[i-1][j-stones[i-1]])
		}
	}

	// 找出最大的j（0 ≤j≤ target）使得dp[n][j]==true，记为maxWeight（即找到不超过sum/2的最大一堆的重量）
	maxWeight := 0
	for j := target; j >= 0; j-- {
		if dp[len(stones)][j] {
			maxWeight = j
			break
		}
	}

	return sum - 2*maxWeight
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
