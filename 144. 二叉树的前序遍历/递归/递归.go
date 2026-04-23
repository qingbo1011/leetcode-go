package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 前序遍历顺序：根节点 → 左子树 → 右子树
	*res = append(*res, node.Val)
	dfs(node.Left, res)
	dfs(node.Right, res)
}

func main() {

}
