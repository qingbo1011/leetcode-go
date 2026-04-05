package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针，单条链表走完了走再从另一条链表头开始，如果相交最后会相遇
	pA, pB := headA, headB
	for pA != pB { // 如果没有相交，最后会是nil==nil
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}

func main() {

}
