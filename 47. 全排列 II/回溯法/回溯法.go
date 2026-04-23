package main

import "sort"

func permuteUnique(nums []int) [][]int {
	// 排序，使相同数字相邻，便于剪枝
	sort.Ints(nums)
	res := [][]int{}
	used := make([]bool, len(nums))
	path := []int{}
	backtrack(nums, used, path, &res)
	return res
}

// backtrack 回溯函数，生成不重复的全排列
func backtrack(nums []int, used []bool, path []int, res *[][]int) {
	// 递归终止条件：当前路径长度等于数组长度，说明得到一个完整排列
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// 遍历所有元素，尝试将未使用的元素加入路径
	for i := 0; i < len(nums); i++ {
		// 剪枝条件：
		// 1. 当前元素已使用，跳过
		if used[i] {
			continue
		}
		// 2. 当前元素与前一个元素相同，且前一个元素未使用，则跳过
		//    原因：在同一层递归中，如果前一个相同的数字没有被使用过，说明它会在后续被尝试，为了避免重复，只取第一个出现的相同数字
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// 做选择：标记该元素已使用，并将其加入路径
		used[i] = true
		path = append(path, nums[i])
		// 递归进入下一层，继续构建剩余部分的排列
		backtrack(nums, used, path, res)
		// 撤销选择（回溯）：移除路径末尾元素，并恢复未使用标记
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {

}
