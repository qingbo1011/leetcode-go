package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 返回二叉树的最大深度（从根到最远叶子节点的节点数）
func maxDepth(root *TreeNode) int {
	// 空树深度为 0
	if root == nil {
		return 0
	}

	depth := 0                 // 记录当前已遍历的层数（深度）
	queue := []*TreeNode{root} // 队列初始化，根节点入队

	// 当队列不为空时，继续处理
	for len(queue) > 0 {
		// 记录当前层的节点数量（此时队列中全部为当前层节点）
		levelSize := len(queue)

		// 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			// 弹出队首节点
			node := queue[0]
			queue = queue[1:]

			// 将当前节点的左右子节点加入队列（作为下一层）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 当前层所有节点处理完毕，深度加 1
		depth++
	}

	return depth
}

func main() {

}
