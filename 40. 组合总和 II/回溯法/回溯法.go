package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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
		// 剪枝：当前数字已大于目标，后面的更大，直接结束循环
		if candidates[i] > target {
			break
		}
		// 去重：如果当前数字和前一个相同，且前一个没有被使用（即在同一层中已经尝试过），跳过
		// 为什么是 i > start && candidates[i] == candidates[i-1]？
		// start是当前层起始索引。当i>start时，candidates[i-1]是本层之前已经处理过的元素。
		// 如果candidates[i] == candidates[i-1]，说明本层已经用相同值做过一次选择，再选会导致重复组合
		// 例如 [1,1,2] 中，两个 1 会生成相同的组合。因此跳过。
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		// 注意：每个数字只能用一次，下一层从i+1开始
		backtrack(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1] // 撤销选择
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
