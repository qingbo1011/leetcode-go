package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	backtrack(nums, 0, path, &res)
	return res
}

// backtrack 回溯枚举所有子集
// start 表示当前可选择的起始索引，保证不重复使用元素
func backtrack(nums []int, start int, path []int, res *[][]int) {
	// 将当前子集加入结果（深拷贝）
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	// 从 start 开始尝试添加每个元素
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])    // 选择当前元素
		backtrack(nums, i+1, path, res) // 递归下一层
		path = path[:len(path)-1]       // 撤销选择
	}
}

func main() {

}
