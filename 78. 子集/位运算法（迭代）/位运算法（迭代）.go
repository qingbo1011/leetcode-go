package main

func subsets(nums []int) [][]int {
	// 对于长度为 n 的数组，子集数量为2ⁿ。
	n := len(nums)
	res := make([][]int, 0, 1<<n)
	// 用一个二进制掩码 mask 从 0 到 2ⁿ-1 遍历，第 i 位为 1 表示选取 nums[i]。
	// 根据掩码构造子集。
	for mask := 0; mask < (1 << n); mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				subset = append(subset, nums[i])
			}
		}
		res = append(res, subset)
	}
	return res
}

func main() {

}
