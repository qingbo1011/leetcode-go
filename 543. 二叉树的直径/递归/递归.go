package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	// 递归计算每个节点的深度，并更新最大直径
	depth(root, &maxDiameter)
	return maxDiameter
}

// depth 返回以 node 为根的子树的最大深度，同时通过指针更新全局最大直径
func depth(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}
	// 左子树深度
	leftDepth := depth(node.Left, maxDiameter)
	// 右子树深度
	rightDepth := depth(node.Right, maxDiameter)

	// 当前节点作为路径转折点时，经过该节点的路径长度为 leftDepth + rightDepth
	// 更新全局最大直径
	*maxDiameter = max(*maxDiameter, leftDepth+rightDepth)

	// 返回当前节点的最大深度（左右子树深度的较大者 + 1）
	return max(leftDepth, rightDepth) + 1
}

func main() {

}
