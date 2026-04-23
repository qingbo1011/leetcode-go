package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		// 一直向左走
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 弹出并访问
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		// 转向右子树
		cur = cur.Right
	}

	return res
}

func main() {

}
