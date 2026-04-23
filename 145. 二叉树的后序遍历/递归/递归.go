package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 后序遍历顺序：左子树 → 右子树 → 根节点
	dfs(node.Left, res)
	dfs(node.Right, res)
	*res = append(*res, node.Val)
}

func main() {

}
