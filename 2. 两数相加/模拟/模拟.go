package main

import "fmt"

func NewListBySlice(arr []int) *ListNode {
	dummyNode := &ListNode{}
	temp := dummyNode
	for _, v := range arr {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	return dummyNode.Next
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := NewListBySlice([]int{2, 4, 3})
	l2 := NewListBySlice([]int{5, 6, 4})
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{} // 注意头结点和首元结点的区别
	temp := head
	carry := 0 // 进位
	// 使用||，一次循环搞定l1，l2不对齐的情况
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		// 在for循环内注意判断l1，l2是否已经到尽头了
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
	if carry > 0 {
		temp.Next = &ListNode{Val: carry}
	}
	return head.Next
}
