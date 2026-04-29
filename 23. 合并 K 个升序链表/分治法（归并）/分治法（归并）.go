package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 调用分治合并函数，合并整个区间 [0, len(lists)-1]
	return merge(lists, 0, len(lists)-1)
}

// 通过分治法递归合并链表数组的指定区间 [left, right]
//
//	lists - 链表数组
//	left  - 区间左边界（包含）
//	right - 区间右边界（包含）
//	返回值：合并后的链表头节点
func merge(lists []*ListNode, left, right int) *ListNode {
	// 递归终止条件：区间内只有一个链表，直接返回该链表
	if left == right {
		return lists[left]
	}
	// 计算中间位置，防止整数溢出
	mid := left + (right-left)/2
	// 递归合并左半部分 [left, mid]
	l1 := merge(lists, left, mid)
	// 递归合并右半部分 [mid+1, right]
	l2 := merge(lists, mid+1, right)
	// 合并两个有序链表
	return mergeTwoLists(l1, l2)
}

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	// 直接接上剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

func main() {

}
