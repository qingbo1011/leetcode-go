package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	tmp := dummyNode
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val { // tmp.Next是要删除的节点
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyNode.Next
}
