package main

func numTrees(n int) int {
	// dp[i]由i个节点组成且节点值从1到i互不相同的二叉搜索树有多少种
	dp := make([]int, n+1)
	dp[0] = 1
	// 转移方程：dp[i] =dp[i] + dp[j-1] * dp[i-j](for i := 2; i <= n; i++)
	for i := 1; i <= n; i++ {
		// j表示以j做二叉搜索树头节点，那么：
		// 左边的节点组成只能是1,2..j-1,一共j-1个节点
		// 右边的节点组成只能是j+1,j+2...i,一共j-1个节点
		// 显然：dp[j] = dp[j-1]*dp[i-j]
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}

	return dp[n]
}

func main() {

}
