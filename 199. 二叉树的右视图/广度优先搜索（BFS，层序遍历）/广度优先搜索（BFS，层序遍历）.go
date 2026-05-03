package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{} // 空树返回空切片
	}
	res := []int{}             // 存储结果
	queue := []*TreeNode{root} // 队列初始化，根节点入队
	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数
		for i := 0; i < levelSize; i++ {
			// 弹出队首节点
			node := queue[0]
			queue = queue[1:]
			// 如果是当前层的最后一个节点，加入结果集
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			// 子节点入队（从左到右，保证层序）
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

func main() {

}
