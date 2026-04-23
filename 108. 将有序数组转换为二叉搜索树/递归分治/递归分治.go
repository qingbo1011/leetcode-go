package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// 由于数组已升序，选择中间元素作为根节点，可保证左右子树节点数平衡。
	// 递归构建左子树（数组左半部分）和右子树（数组右半部分）。
	// 当 left > right 时返回 nil。
	return build(nums, 0, len(nums)-1)
}

// build 递归构建平衡BST
func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 选择中间元素作为根（若长度为偶数，中间偏左或偏右均可，这里取中间偏左）
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)
	return root
}

func main() {

}
