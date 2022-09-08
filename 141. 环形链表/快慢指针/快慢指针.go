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
	slow := head // 慢指针
	fast := head // 快指针
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
