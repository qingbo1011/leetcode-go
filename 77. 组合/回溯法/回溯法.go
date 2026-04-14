package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)
	backtrack(n, k, 1, path, &res)
	return res
}

func backtrack(n int, k int, start int, path []int, res *[][]int) {
	// 回溯法抽象理解：
	// 循环：枚举当前层所有可能的选择（从 start 到上界）。
	// 递归：进入下一层，继续选择下一个数字。
	// 回溯：撤销当前选择，回到上一步状态，尝试其他分支。
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// 剪枝：最多到 n-(k-len(path))+1
	// 为什么循环上界是n-(k-len(path))+1？
	// 因为当前已选了len(path)个数，还需选k-len(path)个数。从i开始，后面至少要有k-len(path)个数字可供选择，即i ≤ n-(k-len(path))+1
	// 例如 n=4, k=2，当 path 为空时，i ≤ 4-2+1=3，所以 i 最大取 3（因为从 3 开始后面还有 4 一个数，刚好够选）。
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		backtrack(n, k, i+1, path, res)
		path = path[:len(path)-1] // 去掉path的最后一个元素
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
