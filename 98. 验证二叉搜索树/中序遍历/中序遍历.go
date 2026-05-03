package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt64
	return inorder(root, &prev)
}

// inorder 中序遍历，prev 指向上一个访问的节点值（指针）
func inorder(node *TreeNode, prev *int) bool {
	if node == nil {
		return true
	}
	// 遍历左子树
	if !inorder(node.Left, prev) {
		return false
	}
	// 检查当前节点是否大于前一个值
	if node.Val <= *prev {
		return false
	}
	// 更新前一个值为当前节点值
	*prev = node.Val
	// 遍历右子树
	return inorder(node.Right, prev)
}

func main() {

}
