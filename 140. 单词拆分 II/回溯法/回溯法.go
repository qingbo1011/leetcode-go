package main

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	memo := make(map[int][]string) // 使用哈希表memo 记录已计算过的起始位置，避免重复计算。
	// 记忆化回溯：
	//  定义递归函数 dfs(start) 返回从下标 start 开始的所有可能的句子（单词列表）。
	//  使用哈希表memo 记录已计算过的起始位置，避免重复计算。
	//  在dfs(start)中，遍历所有可能的单词word（从wordDict中找匹配），
	//  如果s[start:start+len(word)] == word，则递归调用dfs(start+len(word))，并将当前单词与递归结果组合成完整句子。
	// 最终答案：dfs(0)返回所有句子，用空格连接单词。
	return dfs(0, s, wordSet, memo)
}

// dfs 返回从start开始的所有可能句子（单词列表，空格已拼接）
func dfs(start int, s string, wordSet map[string]bool, memo map[int][]string) []string {
	if start == len(s) {
		return []string{""} // 空字符串表示一个合法结尾
	}
	if res, ok := memo[start]; ok {
		return res
	}
	res := []string{}
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if wordSet[word] {
			subSentences := dfs(end, s, wordSet, memo)
			for _, sub := range subSentences {
				if sub == "" {
					res = append(res, word)
				} else {
					res = append(res, word+" "+sub)
				}
			}
		}
	}
	memo[start] = res
	return res
}

func main() {

}
