package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序便于剪枝
	res := make([][]int, 0)
	path := make([]int, 0)
	backtrack(candidates, target, 0, path, &res)
	return res
}

func backtrack(candidates []int, target int, start int, path []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 剪枝：当前数字已大于剩余目标
		if candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		// 注意：下一层递归仍从 i 开始（可以重复使用当前数字）
		backtrack(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1] // 撤销选择
	}
}

func main() {

}
