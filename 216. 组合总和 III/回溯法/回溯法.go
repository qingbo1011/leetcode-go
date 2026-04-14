package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)
	backtrack(k, n, 1, path, &res)
	return res
}

func backtrack(k, n, start int, path []int, res *[][]int) {
	if len(path) == k {
		// 为什么需要检查n == 0？
		// 递归参数n表示剩余需要凑的目标值。
		// 当路径长度达到k时，只有剩余目标恰好为0，才说明当前组合的和等于原始n。
		// 否则（n != 0说明和不够或超过了（但因为剪枝 i > n 阻止了超过，所以n此时只能是正数，表示和不足），不应该记录。
		if n == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*res = append(*res, tmp)
			return
		}
	}
	// 剪枝：最多到9-(k-len(path))+1（k-len(path)表示还需要选多少个数字。）
	// 如果当前和已经大于n，停止继续。
	// 如果剩余可选的数字个数不足k-len(path)，提前返回。
	for i := start; i <= 9-(k-len(path))+1; i++ {
		if i > n { // 当前数字已大于剩余和，后面更大，直接退出
			break
		}
		path = append(path, i)
		backtrack(k, n-i, i+1, path, res)
		path = path[:len(path)-1] // 回溯
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
