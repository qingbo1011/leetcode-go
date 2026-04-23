package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 中序遍历顺序：左子树 → 根节点 → 右子树
	dfs(node.Left, res)
	*res = append(*res, node.Val)
	dfs(node.Right, res)
}

func main() {

}
