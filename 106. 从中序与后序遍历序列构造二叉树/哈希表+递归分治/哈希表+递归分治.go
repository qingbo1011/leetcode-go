package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 根据中序遍历和后序遍历构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 建立中序遍历值到索引的映射，便于 O(1) 查找根节点位置
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}
	// 调用递归辅助函数，区间采用左闭右开 [start, end)
	return buildTreeHelper(inorder, postorder, indexMap, 0, len(inorder), 0, len(postorder))
}

// buildTreeHelper 递归构造二叉树
// inorder: 中序遍历数组
// postorder: 后序遍历数组
// indexMap: 中序遍历值到索引的映射
// inStart, inEnd: 当前子树在中序数组中的区间（左闭右开）
// postStart, postEnd: 当前子树在后序数组中的区间（左闭右开）
// 返回：当前子树的根节点
func buildTreeHelper(inorder, postorder []int, indexMap map[int]int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	// 区间为空，返回 nil
	if inStart >= inEnd || postStart >= postEnd {
		return nil
	}
	// 后序区间的最后一个元素是当前子树的根节点
	rootVal := postorder[postEnd-1]
	root := &TreeNode{Val: rootVal}
	// 在中序中找到根节点的位置
	rootIdx := indexMap[rootVal]
	// 左子树的节点个数
	leftSize := rootIdx - inStart
	// 递归构造左子树
	// 中序区间：[inStart, rootIdx)
	// 后序区间：[postStart, postStart+leftSize)
	root.Left = buildTreeHelper(inorder, postorder, indexMap,
		inStart, rootIdx,
		postStart, postStart+leftSize)
	// 递归构造右子树
	// 中序区间：[rootIdx+1, inEnd)
	// 后序区间：[postStart+leftSize, postEnd-1)  （注意排除最后一个根节点）
	root.Right = buildTreeHelper(inorder, postorder, indexMap,
		rootIdx+1, inEnd,
		postStart+leftSize, postEnd-1)
	return root
}

func main() {

}
