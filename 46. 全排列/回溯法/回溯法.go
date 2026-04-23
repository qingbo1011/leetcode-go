package main

func permute(nums []int) [][]int {
	res := [][]int{}                // 存储所有排列结果
	used := make([]bool, len(nums)) // 标记元素是否已被当前路径使用
	path := []int{}                 // 当前正在构建的排列
	backtrack(nums, used, path, &res)
	return res
}

// backtrack 回溯函数，递归地构建所有排列
// nums: 原始数组
// used: 元素使用标记
// path: 当前已选元素列表
// res: 结果集指针
func backtrack(nums []int, used []bool, path []int, res *[][]int) {
	// 递归终止条件：当前路径长度等于数组长度，说明得到一个完整排列
	if len(path) == len(nums) {
		// 拷贝当前路径，因为后续回溯会修改 path 切片
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	// 遍历所有元素，尝试将未使用的元素加入路径
	for i := 0; i < len(nums); i++ {
		// 剪枝：跳过已经使用过的元素
		if used[i] {
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
