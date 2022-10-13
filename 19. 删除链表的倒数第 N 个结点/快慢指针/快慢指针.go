package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for i := 0; i < n; i++ { // 双指针就位
		p2 = p2.Next
	}
	for p2.Next != nil { // p1就位到要删除节点的前一个位置
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next // 删除节点
	return dummy.Next
}
