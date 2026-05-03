package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	cur := root // 当前处理的节点
	for cur != nil {
		// 如果当前节点有左子树，则需要将左子树插入到 cur 和 cur.Right 之间
		if cur.Left != nil {
			// 1. 找到左子树的最右节点
			leftMostRight := cur.Left
			for leftMostRight.Right != nil {
				leftMostRight = leftMostRight.Right
			}
			// 2. 将 cur 的右子树接到左子树最右节点的右侧
			leftMostRight.Right = cur.Right
			// 3. 将左子树整体移动到右指针位置，并清空左指针
			cur.Right = cur.Left
			cur.Left = nil
		}
		// 4. 继续处理下一个节点（即原右子树或新移过来的右子树）
		cur = cur.Right
	}
}

func main() {

}
