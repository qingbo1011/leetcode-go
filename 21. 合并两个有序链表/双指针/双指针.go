package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	// 直接接上剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

func main() {

}
