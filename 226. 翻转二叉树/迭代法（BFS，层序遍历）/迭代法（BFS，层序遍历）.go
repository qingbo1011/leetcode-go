package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// 空树直接返回
	if root == nil {
		return nil
	}
	// 队列初始化，根节点入队
	queue := []*TreeNode{root}
	// 当队列不为空时，继续处理
	for len(queue) > 0 {
		// 弹出队首节点
		node := queue[0]
		queue = queue[1:]
		// 交换当前节点的左右子树（核心翻转操作）
		node.Left, node.Right = node.Right, node.Left
		// 将非空左子节点入队，以便后续翻转其子树
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		// 将非空右子节点入队
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	// 返回根节点（原树结构已原地翻转）
	return root
}

func main() {

}
