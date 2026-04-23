package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// 空树视为对称
	if root == nil {
		return true
	}

	// 初始化队列，将根节点的左右子节点成对入队
	queue := []*TreeNode{root.Left, root.Right}

	// 当队列不为空时，持续比较成对节点
	for len(queue) > 0 {
		// 弹出队首的两个节点（成对比较）
		u := queue[0]
		v := queue[1]
		queue = queue[2:] // 移除已弹出的两个节点

		// 情况1：两个节点都为空，说明这一对符合对称，继续比较下一对
		if u == nil && v == nil {
			continue
		}
		// 情况2：其中一个为空，或值不相等，则不对称
		if u == nil || v == nil || u.Val != v.Val {
			return false
		}

		// 将下一层需要比较的节点成对入队（注意顺序：镜像位置）
		// 左子树的左孩子 与 右子树的右孩子 成对
		// 左子树的右孩子 与 右子树的左孩子 成对
		queue = append(queue, u.Left, v.Right, u.Right, v.Left)
	}

	// 所有成对节点均比较通过，树对称
	return true
}

func main() {

}
