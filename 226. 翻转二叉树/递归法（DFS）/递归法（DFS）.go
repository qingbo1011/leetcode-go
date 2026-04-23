package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(node *TreeNode) {
	if node == nil {
		return
	}
	// 交换左右子树
	node.Left, node.Right = node.Right, node.Left
	// 递归翻转左右子树
	invert(node.Left)
	invert(node.Right)
}

func main() {

}
