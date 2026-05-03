package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 建立中序遍历值到索引的映射，便于快速定位根节点
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}
	// 调用递归辅助函数，区间采用左闭右开 [start, end)
	return buildTreeHelper(preorder, indexMap, 0, len(preorder), 0, len(inorder))
}

// buildTreeHelper 递归构造二叉树
// preorder: 前序遍历数组
// indexMap: 中序遍历值到索引的映射
// preStart, preEnd: 当前子树在前序数组中的区间（左闭右开）
// inStart, inEnd: 当前子树在中序数组中的区间（左闭右开）
// 返回：当前子树的根节点
func buildTreeHelper(preorder []int, indexMap map[int]int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	// 区间为空，返回 nil
	if preStart >= preEnd || inStart >= inEnd {
		return nil
	}
	// 前序区间的第一个元素就是当前子树的根节点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	// 在中序中找到根节点的位置
	inRootIdx := indexMap[rootVal]
	// 左子树的节点个数 = 根节点在中序中的位置 - 中序区间左边界
	leftSize := inRootIdx - inStart
	// 递归构造左子树
	// 前序区间：根节点之后 leftSize 个元素属于左子树
	// 中序区间：根节点左侧部分
	root.Left = buildTreeHelper(preorder, indexMap,
		preStart+1, preStart+1+leftSize,
		inStart, inRootIdx)
	// 递归构造右子树
	// 前序区间：左子树之后到 preEnd 的部分属于右子树
	// 中序区间：根节点右侧部分
	root.Right = buildTreeHelper(preorder, indexMap,
		preStart+1+leftSize, preEnd,
		inRootIdx+1, inEnd)
	return root
}

func main() {

}
