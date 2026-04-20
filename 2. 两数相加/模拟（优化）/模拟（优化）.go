package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	carry := 0 // 进位
	// 使用||，一次循环搞定l1，l2不对齐的情况
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10 // 得出当前位的结果，以及进位
		temp.Next = &ListNode{Val: sum}
		temp = temp.Next
	}
	// 别忘了最后进位的情况
	if carry > 0 {
		temp.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}

func main() {

}
