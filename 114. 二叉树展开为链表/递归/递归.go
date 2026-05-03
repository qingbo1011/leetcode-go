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
	// 递归展开左子树和右子树
	flatten(root.Left)
	flatten(root.Right)
	// 保存当前右子树
	right := root.Right
	// 将左子树移到右边，然后左子树置空
	root.Right = root.Left
	root.Left = nil
	// 找到当前右子树的末尾（即原先左子树的末尾）
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	// 将原先的右子树接到末尾
	cur.Right = right
}

func main() {

}
