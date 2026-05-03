package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0 // 记录已经访问过的节点数
	ans := 0   // 存储第 k 小的值
	inorder(root, &count, k, &ans)
	return ans
}

// inorder 递归中序遍历，当找到第 k 个节点时记录答案并提前终止
// node - 当前节点
// count - 指向计数的指针，每访问一个节点就加1
// k - 目标序号
// ans - 指向答案的指针
func inorder(node *TreeNode, count *int, k int, ans *int) {
	if node == nil {
		return
	}
	// 先遍历左子树
	inorder(node.Left, count, k, ans)
	// 访问当前节点：计数加1
	*count++
	if *count == k {
		*ans = node.Val
		return // 提前返回，不再继续遍历
	}
	// 再遍历右子树（只有未找到时才需要遍历）
	inorder(node.Right, count, k, ans)
}

func main() {

}
