package main

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		maxLen = max(maxLen, len(word))
		wordSet[word] = true
	}
	// 状态定义：dp[i]表示字符串s的前i个字符（即 s[0:i]）能否被成功拆分成字典中的单词
	dp := make([]bool, len(s)+1)
	// 初始化：dp[0] = true，表示空字符串可以被拆分（不选任何单词）
	dp[0] = true
	// 状态转移方程：对于每个位置i（1 ≤ i ≤ n），遍历所有可能的分割点j（0 ≤ j < i）：
	// 如果 dp[j] == true 且子串 s[j:i] 存在于字典 wordDict 中，则 dp[i] = true
	// 优化：子串 s[j:i] 存在于字典 wordDict 中，那么子串s[j:i] 的长度一定不能超过maxLen，即i-j<=maxLen，将时间复杂度降为 O(n * maxLen)
	for i := 0; i <= len(s); i++ {
		// 直接检查靠近i的j（即较短的后缀子串），符合直觉：我们只关心长度不超过maxLen的分割点。
		for j := i - 1; j >= 0 && i-j <= maxLen; j-- {
			dp[i] = dp[j] && wordSet[s[j:i]]
			// 记录dp[i]结果退出循环
			if dp[i] == true {
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {

}
