package main

func climbStairs(n int, costs []int) int {
	// 滚动数组：只保存最近3个台阶的状态，因为从第 i 级只能由 i-1, i-2, i-3 转移过来
	// dp[i%3] 存储到达第 i 级台阶的最小成本
	dp := make([]int, 3) // 初始化为0，dp[0]对应i=0时为0

	// 从第1级台阶开始计算，直到第n级
	for i := 1; i <= n; i++ {
		// cur 表示到达第 i 级的最小成本，初始化为最大整数（方便取最小值）
		cur := int(^uint(0) >> 1) // 在Go中，^uint(0)得到全1，右移1位得到最大有符号正整数

		// 尝试三种可能的步长：1步、2步、3步
		for step := 1; step <= 3; step++ {
			// 前驱台阶编号为 i-step，必须非负才有效
			if i-step >= 0 {
				// 从滚动数组中获取到达前驱台阶的最小成本
				prev := dp[(i-step)%3]

				// 计算从 i-step 跳到 i 的总成本：
				//   prev 是到达前驱的最小成本
				//   costs[i-1] 是目标台阶 i 的成本（因为 costs 下标0对应第1级）
				//   step*step 是 (j-i)² 的平方项
				cost := prev + costs[i-1] + step*step

				// 取最小值
				if cost < cur {
					cur = cost
				}
			}
		}

		// 将计算出的第 i 级的最小成本存入滚动数组的对应位置（使用 i%3 循环覆盖）
		dp[i%3] = cur
	}

	// 返回第 n 级的最小成本，同样通过取模定位
	return dp[n%3]
}

func main() {

}
