package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	dfs(root, 0, &res)
	return res
}

// dfs 递归遍历，depth 表示当前深度
// res 指针用于在递归中修改结果切片
func dfs(node *TreeNode, depth int, res *[]int) {
	if node == nil {
		return
	}
	// 如果当前深度还没有记录，说明这是该层第一个被访问的节点（即最右侧节点）
	if depth == len(*res) {
		*res = append(*res, node.Val)
	}
	// 优先递归右子树，保证右侧节点先被访问
	dfs(node.Right, depth+1, res)
	dfs(node.Left, depth+1, res)
}

func main() {

}
