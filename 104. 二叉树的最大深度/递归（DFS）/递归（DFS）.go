package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return depth(root)
}

func depth(node *TreeNode) int {
	// 如果根节点为空，深度为0
	if node == nil {
		return 0
	}
	// 最大深度 = 1 + max(左子树最大深度, 右子树最大深度)
	return 1 + max(depth(node.Left), depth(node.Right))
}

func main() {

}
