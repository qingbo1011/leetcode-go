package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 返回二叉树的层序遍历结果（利用 DFS 递归实现）
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}   // 存储最终结果，每一层是一个子切片
	dfs(root, 0, &res) // 从根节点开始，深度为 0
	return res
}

// dfs 递归遍历二叉树，按深度将节点值存入对应层
func dfs(node *TreeNode, depth int, res *[][]int) {
	if node == nil {
		return
	}
	// 如果当前深度对应的层还未创建，则先添加一个空切片
	if depth == len(*res) {
		*res = append(*res, []int{})
	}
	// 将当前节点值加入对应深度层的切片中
	(*res)[depth] = append((*res)[depth], node.Val)
	// 递归处理左子树，深度+1
	dfs(node.Left, depth+1, res)
	// 递归处理右子树，深度+1
	dfs(node.Right, depth+1, res)
}

func main() {

}
