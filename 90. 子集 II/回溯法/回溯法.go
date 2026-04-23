package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 排序，使重复元素相邻
	res := [][]int{}
	path := []int{}
	backtrack(nums, 0, path, &res)
	return res
}

// backtrack 回溯枚举所有子集（含去重）
// start 表示当前可选择的起始索引
func backtrack(nums []int, start int, path []int, res *[][]int) {
	// 将当前子集加入结果（深拷贝）
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	// 从 start 开始尝试添加每个元素
	for i := start; i < len(nums); i++ {
		// 剪枝：如果当前元素与前一个元素相同，且前一个元素在同一层已经被尝试过，则跳过
		// 注意：i > start 保证只跳过同一层中重复的选项，不影响不同层的重复元素
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrack(nums, i+1, path, res)
		path = path[:len(path)-1] // 撤销选择
	}
}

func main() {

}
