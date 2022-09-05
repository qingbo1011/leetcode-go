package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head      // 慢指针
	p2 := head.Next // 快指针
	for p2 != nil && p2.Next != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}
