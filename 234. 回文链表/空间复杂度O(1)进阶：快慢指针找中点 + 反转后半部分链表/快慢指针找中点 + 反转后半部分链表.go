package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 1.快慢指针找中点
	slow, fast := head, head
	// 保证慢指针停在前半部分的最后一个节点，方便反转后半部分。
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 2.反转后半部分
	secondHalf := reverseList(slow.Next)
	// 3.比较前后两部分
	p1, p2 := head, secondHalf
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// 4.恢复链表（可选）
	slow.Next = reverseList(secondHalf)

	return true
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

func main() {

}
