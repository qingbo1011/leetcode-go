package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 使用int64避免边界溢出
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// dfs 检查以 root 为根的子树是否满足 BST 条件，且所有节点值在 (min, max) 范围内
func dfs(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	val := int64(root.Val)
	// 当前节点值必须在 (min, max) 开区间内
	if val <= min || val >= max {
		return false
	}
	// 左子树：最大值更新为当前节点值
	// 右子树：最小值更新为当前节点值
	return dfs(root.Left, min, val) && dfs(root.Right, val, max)
}

func main() {

}
