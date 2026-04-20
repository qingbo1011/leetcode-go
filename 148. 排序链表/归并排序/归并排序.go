package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 1. 计算链表总长度
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	// 哑节点，方便处理头节点变化
	dummy := &ListNode{Next: head}
	// 2. 步长 step 表示当前要合并的子链表长度，从 1 开始，每次翻倍
	for step := 1; step < length; step <<= 1 {
		prev := dummy     // prev 指向已经合并好的链表的尾部
		cur := dummy.Next // cur 指向当前待处理的子链表头部
		// 3. 遍历整个链表，按步长分割并两两合并
		for cur != nil {
			// 左子链表：从 cur 开始，长度为 step
			left := cur
			// 切掉前 step 个节点，返回剩余部分的头（即右子链表的头）
			right := cut(left, step)
			// 再从右子链表头切掉 step 个节点，返回剩余部分的头（下一组的开始）
			nextStart := cut(right, step)
			// 合并 left 和 right 这两个有序子链表
			prev.Next = merge(left, right)
			// 移动 prev 到合并后链表的尾部，以便连接下一组
			for prev.Next != nil {
				prev = prev.Next
			}
			// 继续处理下一组
			cur = nextStart
		}
	}

	return dummy.Next
}

// cut 切断链表前n个节点，返回剩余部分的头节点
// 例如：链表 1->2->3->4，cut(1,2) 后，链表变为 1->2，返回 3->4
func cut(head *ListNode, n int) *ListNode {
	cur := head
	// 这里是n > 1，因为要切断链表前n个节点，因此需要cur.Next = nil
	for cur != nil && n > 1 {
		cur = cur.Next
		n--
	}
	if cur == nil {
		return nil
	}
	next := cur.Next
	cur.Next = nil
	return next
}

// merge 合并两个有序链表(参考leetcode：21. 合并两个有序链表)
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

func main() {

}
