package main

func uniquePaths(m int, n int) int {
	// 计算 C(m+n-2, m-1)
	ans := 1
	total := m + n - 2
	k := min(m-1, n-1) // 选择较小的组合数计算
	for i := 1; i <= k; i++ {
		ans = ans * (total - k + i) / i
	}
	return ans
}

func main() {

}
