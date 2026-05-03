package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{} // 模拟递归栈
	cur := root            // 当前节点
	count := 0             // 已访问节点数

	// 当还有节点未处理（当前节点非空 或 栈非空）时继续
	for cur != nil || len(stack) > 0 {
		// 1. 一直向左走，将路径上的节点入栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 2. 弹出栈顶节点（最左下方未访问节点）
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 3. 访问当前节点
		count++
		if count == k {
			return cur.Val
		}

		// 4. 转向右子树（重复上述过程）
		cur = cur.Right
	}
	return -1 // 理论上不会执行到这里，因为 1 <= k <= n
}

func main() {

}
