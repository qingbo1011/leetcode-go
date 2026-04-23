package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	// 边界条件：空树直接返回空切片
	if root == nil {
		return []int{}
	}

	res := []int{}              // 最终结果切片
	stack1 := []*TreeNode{root} // 栈1：用于前序遍历的变体（根 -> 右 -> 左）
	stack2 := []*TreeNode{}     // 栈2：用于逆序存储访问顺序（最终得到后序）

	// 第一阶段：利用栈1模拟“根 -> 右 -> 左”的遍历顺序，结果压入栈2
	for len(stack1) > 0 {
		// 弹出栈1顶部节点
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		// 将当前节点压入栈2（后续会逆序输出，实现左右根的顺序）
		stack2 = append(stack2, node)

		// 注意：先压左孩子，再压右孩子，保证栈1弹出的顺序是“根 -> 右 -> 左”
		// 因为栈是后进先出，所以先压左，则左后出；后压右，则右先出。
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	// 第二阶段：栈2中存储的顺序是“根 -> 右 -> 左”的逆序，即“左 -> 右 -> 根”
	// 从栈顶（最后一个元素）开始逆序弹出，得到后序遍历结果
	for i := len(stack2) - 1; i >= 0; i-- {
		res = append(res, stack2[i].Val)
	}

	return res
}

func main() {

}
