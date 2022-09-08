package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		tmp.Next = list1
		tmp = tmp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		tmp.Next = list2
		tmp = tmp.Next
		list2 = list2.Next
	}
	return res.Next
}
