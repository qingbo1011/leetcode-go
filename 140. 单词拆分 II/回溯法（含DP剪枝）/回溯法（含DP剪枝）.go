package main

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	// 1. DP 预处理可达性
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	// 如果整体不可达，直接返回空
	if !dp[n] {
		return []string{}
	}
	memo := make(map[int][]string)
	return dfsWithDP(0, s, wordSet, dp, memo)
}

func dfsWithDP(start int, s string, wordSet map[string]bool, dp []bool, memo map[int][]string) []string {
	if start == len(s) {
		return []string{""}
	}
	if res, ok := memo[start]; ok {
		return res
	}
	res := []string{}
	for end := start + 1; end <= len(s); end++ {
		if dp[end] && wordSet[s[start:end]] { // 只有可达且单词匹配才继续
			sub := dfsWithDP(end, s, wordSet, dp, memo)
			for _, sentence := range sub {
				if sentence == "" {
					res = append(res, s[start:end])
				} else {
					res = append(res, s[start:end]+" "+sentence)
				}
			}
		}
	}
	memo[start] = res
	return res
}

func main() {

}
