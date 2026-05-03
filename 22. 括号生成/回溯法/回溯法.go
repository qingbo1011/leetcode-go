package main

func generateParenthesis(n int) []string {
	res := []string{}
	// ● 递归终止条件：当 left == right == n 时，得到一个有效组合。
	// ● 剪枝条件（保证有效性）：
	//  ○ 如果 left < n，可以放置一个左括号。
	//  ○ 如果 right < left，可以放置一个右括号。
	backtrack(n, 0, 0, "", &res)
	return res
}

// backtrack 递归生成括号组合
// n: 括号对数
// left: 已使用的左括号数
// right: 已使用的右括号数
// cur: 当前构建的字符串
// res: 结果集指针
func backtrack(n, left, right int, cur string, res *[]string) {
	if left == n && right == n {
		// 左右括号都已用完，得到一个有效组合
		*res = append(*res, cur)
		return
	}
	// 1. 放置左括号（如果尚未用完）
	if left < n {
		backtrack(n, left+1, right, cur+"(", res)
	}
	// 2. 放置右括号（仅当右括号数小于左括号数时，才能保证有效）
	if right < left {
		backtrack(n, left, right+1, cur+")", res)
	}
}

func main() {

}
