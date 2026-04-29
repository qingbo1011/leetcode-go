package main

import "slices"

// 预处理交替排列的方案数
// f 数组存储当剩余位置数为 i 时（奇偶模式固定）的排列数，即 (剩余奇数个数)! * (剩余偶数个数)!
// f[0] = 1, f[1] = 1, f[2] = 1, f[3] = 2, f[4] = 2, f[5] = 6, f[6] = 6, ...
var f = []int{1}

func init() {
	// 构造 f 数组，直到乘积超过 1e15（因为题目中的 k ≤ 1e15）
	for i := 1; f[len(f)-1] < 1e15; i++ {
		// 每次追加两次 f[last]*i，使得 f[2m] = m! * m!，f[2m+1] = m! * (m+1)!
		f = append(f, f[len(f)-1]*i)
		f = append(f, f[len(f)-1]*i)
	}
}

func permute(n int, K int64) []int {
	// 将 k 转换为 0‑based 索引，方便整除和取余操作
	k := int(K - 1)

	// 边界检查：如果 n 在预计算范围内，且 k 超过总排列数，则返回空切片
	// 总排列数 = f[n] * (2 - n%2)，当 n 为偶数时乘2（两种奇偶开头），奇数时乘1
	if n < len(f) && k >= f[n]*(2-n%2) {
		return nil
	}

	// cand 的两个子切片分别存储当前可用的偶数和奇数，均保持升序
	// cand[0] 存放偶数，cand[1] 存放奇数
	cand := [2][]int{}
	// 升序收集所有偶数
	for i := 2; i <= n; i += 2 {
		cand[0] = append(cand[0], i)
	}
	// 升序收集所有奇数
	for i := 1; i <= n; i += 2 {
		cand[1] = append(cand[1], i)
	}

	ans := make([]int, n) // 结果数组
	parity := 1           // 当前需要填入的数的奇偶性：1 表示奇数，0 表示偶数
	// 逐位确定排列中的每个数字
	for i := 0; i < n; i++ {
		rem := n - 1 - i // 剩余待填位置数
		j := 0           // 当前候选列表中的索引
		if rem < len(f) {
			// 每个候选数字对应的排列数 size = f[rem]
			size := f[rem]
			// 确定当前位应该选择候选列表中的第 j 个（0‑based）
			j = k / size
			k %= size // 更新剩余 k 用于后续位

			// 特殊情况：n 为偶数且当前是第一个位置（i==0）
			// 第一个位置既可以是奇数又可以是偶数，需要根据 j 决定 parity
			if n%2 == 0 && i == 0 {
				// j 的奇偶性决定了第一个数的奇偶性：0→偶数，1→奇数
				parity = 1 - j%2
				// 因为每种奇偶性下第一个数字的候选列表长度相等，所以 j 需除以2
				j /= 2
			}
		}
		// 从对应奇偶性的候选列表中取出第 j 个数字
		ans[i] = cand[parity][j]
		// 删除已使用的数字，保持候选列表顺序
		cand[parity] = slices.Delete(cand[parity], j, j+1)
		// 下一个位置的奇偶性必须与当前相反
		parity ^= 1
	}
	return ans
}

func main() {

}
