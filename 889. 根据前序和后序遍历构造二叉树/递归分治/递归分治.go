package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// 建立后序遍历值到索引的映射，便于快速定位
	postIndexMap := make(map[int]int)
	for i, val := range postorder {
		postIndexMap[val] = i
	}
	return build(preorder, 0, len(preorder), postorder, 0, len(postorder), postIndexMap)
}

// build 递归构造二叉树
// preorder: 前序遍历数组
// preStart, preEnd: 当前子树在前序数组中的区间（左闭右开）
// postorder: 后序遍历数组
// postStart, postEnd: 当前子树在后序数组中的区间（左闭右开）
// postIndexMap: 后序遍历值到索引的映射
// 返回：当前子树的根节点
func build(preorder []int, preStart, preEnd int,
	postorder []int, postStart, postEnd int,
	postIndexMap map[int]int) *TreeNode {

	if preStart >= preEnd {
		return nil
	}
	// 根节点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	// 如果当前区间只有一个节点，直接返回
	if preEnd-preStart == 1 {
		return root
	}
	// 左子树的根节点值（前序中根的下一个）
	leftRootVal := preorder[preStart+1]
	// 在后序中找到左子树根的位置
	leftRootIdx := postIndexMap[leftRootVal]
	// 左子树的节点个数
	leftSize := leftRootIdx - postStart + 1
	// 递归构建左子树
	root.Left = build(preorder, preStart+1, preStart+1+leftSize,
		postorder, postStart, postStart+leftSize,
		postIndexMap)
	// 递归构建右子树
	root.Right = build(preorder, preStart+1+leftSize, preEnd,
		postorder, postStart+leftSize, postEnd-1, // 排除最后一个根节点
		postIndexMap)
	return root
}

func main() {

}
