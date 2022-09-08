package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	pre := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 先移动在判断，因为slow和fast初始都在head上
		if slow == fast {
			for slow != pre {
				slow = slow.Next
				pre = pre.Next
			}
			return pre
		}
	}
	return nil
}
