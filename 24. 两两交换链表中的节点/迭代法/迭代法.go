package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		first, second := pre.Next, pre.Next.Next
		// 交换
		pre.Next = second
		first.Next = second.Next
		second.Next = first
		// 移动pre到下一对的前驱
		pre = first
	}

	return dummy.Next
}

func main() {

}
