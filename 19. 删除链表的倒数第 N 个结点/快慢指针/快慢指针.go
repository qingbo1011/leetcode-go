package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	// p2向前走n步,双指针就位
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	// 当p2.Next == nil时，p1.Next就是要删除的节点
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next

	return dummy.Next
}

func main() {

}
