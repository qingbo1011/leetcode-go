package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 初始化栈，将根节点入栈
	stack := []*TreeNode{root}
	// prev 记录上一次处理的节点，用于连接
	var prev *TreeNode

	// 栈不为空时继续处理
	for len(stack) > 0 {
		// 弹出栈顶节点（当前要处理的节点）
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果 prev 不为空，说明上一个节点需要连接到当前节点
		if prev != nil {
			prev.Right = cur // 右指针指向当前节点
			prev.Left = nil  // 左指针置空（保证链表形态）
		}

		// 由于栈是后进先出（LIFO），为了确保先序遍历的“左-右”顺序，
		// 需要先将右子节点入栈，再将左子节点入栈，这样左子节点会先被弹出处理。
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

		// 更新 prev 为当前节点，用于下一轮连接
		prev = cur
	}
}

func main() {

}
