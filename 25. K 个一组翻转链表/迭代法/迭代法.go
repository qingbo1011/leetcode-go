package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		// 检查是否有k个节点
		end := pre
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// 翻转[pre.Next, end]区间
		start := pre.Next
		nextStart := end.Next
		// 翻转子链表
		var preNode *ListNode
		cur := start
		for cur != nextStart {
			tmp := cur.Next
			cur.Next = preNode
			preNode = cur
			cur = tmp
		}
		// 将翻转后的子链表接回
		pre.Next = end
		start.Next = nextStart
		// 移动pre到下一组的前一个节点
		pre = start
	}
	return dummy.Next
}

func main() {

}
