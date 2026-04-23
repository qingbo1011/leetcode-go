package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}           // 存储最终结果，每一层是一个子切片
	queue := []*TreeNode{root} // 初始化队列，根节点入队

	// 当队列不为空时，继续处理
	for len(queue) > 0 {
		// 记录当前层的节点数（此时队列中的节点全部属于当前层）
		levelSize := len(queue)
		level := []int{} // 存储当前层的节点值
		// 依次处理当前层的每个节点
		for i := 0; i < levelSize; i++ {
			// 弹出队首节点
			node := queue[0]
			queue = queue[1:]
			// 记录当前节点的值
			level = append(level, node.Val)

			// 将左右子节点加入队列（作为下一层）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将当前层的结果加入最终结果集
		res = append(res, level)
	}

	return res
}

func main() {

}
