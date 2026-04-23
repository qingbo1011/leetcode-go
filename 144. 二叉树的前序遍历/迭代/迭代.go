package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// 弹出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 记录根节点
		res = append(res, node.Val)
		// 栈先进后出，所以先压右子树，再压左子树
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

func main() {

}
