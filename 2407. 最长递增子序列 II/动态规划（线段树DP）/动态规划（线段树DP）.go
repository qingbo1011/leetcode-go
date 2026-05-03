package main

// SegTree 线段树结构体，用于维护区间最大值
type SegTree struct {
	tree []int // 线段树数组，下标从1开始，大小为4*n
	n    int   // 值域范围（1 ~ maxVal）
}

// NewSegTree 构造函数，创建一个线段树实例
// maxVal: 数值的最大范围（例如数组中元素的最大值）
func NewSegTree(maxVal int) *SegTree {
	// 分配4倍空间，防止越界，同时多分配5个冗余
	return &SegTree{
		tree: make([]int, 4*maxVal+5),
		n:    maxVal,
	}
}

// Update 单点更新：将位置 idx 的值更新为 max(原值, val)
func (st *SegTree) Update(idx, val int) {
	st.update(1, 1, st.n, idx, val)
}

// 递归更新内部方法
// node: 当前节点编号
// l, r: 当前节点代表的区间范围 [l, r]
// idx: 需要更新的叶子节点位置
// val: 新的值
func (st *SegTree) update(node, l, r, idx, val int) {
	if l == r {
		// 叶子节点，直接比较并取较大值（因为同一个数值可能多次出现）
		if val > st.tree[node] {
			st.tree[node] = val
		}
		return
	}
	mid := (l + r) / 2
	if idx <= mid {
		st.update(node*2, l, mid, idx, val)
	} else {
		st.update(node*2+1, mid+1, r, idx, val)
	}
	// 回溯时更新父节点为左右孩子的最大值
	st.tree[node] = max(st.tree[node*2], st.tree[node*2+1])
}

// Query 区间查询最大值
// ql, qr: 查询区间 [ql, qr]
// 注意如果 ql > qr，返回 0
func (st *SegTree) Query(ql, qr int) int {
	if ql > qr {
		return 0
	}
	return st.query(1, 1, st.n, ql, qr)
}

// 递归查询内部方法
func (st *SegTree) query(node, l, r, ql, qr int) int {
	// 当前区间完全被查询区间覆盖，直接返回节点值
	if ql <= l && r <= qr {
		return st.tree[node]
	}
	mid := (l + r) / 2
	res := 0
	// 左区间有重叠部分
	if ql <= mid {
		res = max(res, st.query(node*2, l, mid, ql, qr))
	}
	// 右区间有重叠部分
	if qr > mid {
		res = max(res, st.query(node*2+1, mid+1, r, ql, qr))
	}
	return res
}

// 计算最长递增子序列，且相邻元素差值不超过 k
func lengthOfLIS(nums []int, k int) int {
	// 动态规划思路：定义dp[val]表示以数值val结尾的最长合法子序列长度。
	// 遍历数组nums，对于当前数x，我们需要找到所有可以拼接在x前面的数值y，
	// 满足 y < x 且 x - y <= k（即 y 在 [x-k, x-1] 范围内），
	// 并取这些 dp[y] 的最大值，然后 dp[x] = max(dp[x], best + 1)。
	// 1. 找到数组中的最大值，确定值域范围
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// 2. 初始化线段树，所有dp初始值为0
	seg := NewSegTree(maxVal)
	ans := 0 // 全局最长长度
	// 3. 遍历数组中的每个数
	for _, x := range nums {
		// 可转移的前驱数值范围：[x-k, x-1]
		left := x - k
		if left < 1 {
			left = 1
		}
		right := x - 1
		// 查询区间 [left, right] 内的最大 dp 值
		best := seg.Query(left, right)
		// 当前以 x 结尾的最长合法子序列长度
		cur := best + 1
		// 更新全局答案
		if cur > ans {
			ans = cur
		}
		// 更新线段树中位置 x 的值为 max(原值, cur)
		seg.Update(x, cur)
	}
	return ans
}

func main() {

}
