package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	// 状态定义：对于每个节点root，定义：
	// dfs(root) 返回一个长度为 2 的数组 [偷该节点的最大收益, 不偷该节点的最大收益]
	val := dfs(root)
	// 最终答案：max(偷根节点, 不偷根节点)
	return max(val[0], val[1])
}

// dfs 返回 [偷当前节点的最大值, 不偷当前节点的最大值]
func dfs(node *TreeNode) [2]int {
	// 边界条件：空节点返回 [0, 0]
	if node == nil {
		return [2]int{0, 0}
	}
	// 状态转移：
	// 偷当前节点：则不能偷左右子节点，收益 = root.Val + 左子不偷 + 右子不偷
	// 不偷当前节点：则可以偷或不偷左右子节点，取最大值，收益 = max(左子偷, 左子不偷) + max(右子偷, 右子不偷)
	left := dfs(node.Left)   // 得到left[0]左子偷最大值，left[1]左子不偷最大值
	right := dfs(node.Right) // 得到right[0]右子偷最大值，right[1]右子不偷最大值
	// 偷当前节点
	robCurr := node.Val + left[1] + right[1]
	// 不偷当前节点
	notRobCurr := max(left[0], left[1]) + max(right[0], right[1])
	// 返回 [偷当前节点的最大值, 不偷当前节点的最大值]
	return [2]int{robCurr, notRobCurr}
}

func main() {

}
