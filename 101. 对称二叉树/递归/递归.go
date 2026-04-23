package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// isMirror 判断两棵树是否互为镜像
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	// 两棵树互为镜像的条件：
	//  a. 根节点值相等。
	//  b. 左树的左子树 与 右树的右子树 互为镜像。
	//  c. 左树的右子树 与 右树的左子树 互为镜像。
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func main() {

}
