package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 保存下一个节点
		cur.Next = pre   // 反转当前节点的指针
		pre = cur        // 移动 prev
		cur = next       // 移动 curr

	}
	return pre
}
