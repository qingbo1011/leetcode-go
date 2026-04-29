package main

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	// 计算原始利润（不进行任何修改）
	var base int64 = 0
	for i := 0; i < n; i++ {
		// strategy[i] * prices[i] 即为第 i 天的贡献（买入为负，卖为正，持有为0）
		base += int64(strategy[i] * prices[i])
	}

	// 构造两个辅助数组，用于快速计算修改后的增量
	// A[i]：若第 i 天位于修改区间的“前半段”（需变为持有 0），则对该天的利润增量 = 0 - 原贡献 = -strategy[i]*prices[i]
	// B[i]：若第 i 天位于修改区间的“后半段”（需变为卖出 1），则对该天的利润增量 = 1*prices[i] - 原贡献 = prices[i]*(1 - strategy[i])
	A := make([]int64, n)
	B := make([]int64, n)
	for i := 0; i < n; i++ {
		A[i] = int64(-strategy[i] * prices[i])
		B[i] = int64(prices[i] * (1 - strategy[i]))
	}

	// 构建 A 和 B 的前缀和，方便快速计算任意连续子数组的和
	preA := make([]int64, n+1)
	preB := make([]int64, n+1)
	for i := 0; i < n; i++ {
		preA[i+1] = preA[i] + A[i]
		preB[i+1] = preB[i] + B[i]
	}

	half := k / 2        // 前半段的长度（k 一定是偶数）
	var maxInc int64 = 0 // 记录最大增量（可以不修改，因此增量至少为0）

	// 遍历所有可能的修改区间起始位置 i
	// 区间长度为 k，起始位置 i 的范围是 [0, n-k]
	for i := 0; i <= n-k; i++ {
		// 前半段区间 [i, i+half-1]
		leftStart := i
		leftEnd := i + half - 1
		// 后半段区间 [i+half, i+k-1]
		rightStart := i + half
		rightEnd := i + k - 1

		// 计算该修改方案的总增量 = 前半段增量之和 + 后半段增量之和
		inc := (preA[leftEnd+1] - preA[leftStart]) + (preB[rightEnd+1] - preB[rightStart])

		// 更新最大增量
		if inc > maxInc {
			maxInc = inc
		}
	}

	// 最大利润 = 原始利润 + 最大增量（如果增量小于0，则取0，因为我们可以不修改）
	return base + maxInc
}

func main() {

}
