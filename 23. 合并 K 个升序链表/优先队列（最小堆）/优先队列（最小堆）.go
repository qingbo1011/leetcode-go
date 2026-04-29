package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 将所有链表的头节点放入最小堆（按节点值排序）。
	// 弹出堆顶节点，加入结果链表，并将该节点的下一个节点（若存在）入堆。
	// 重复直到堆空。

	// 初始化一个空的最小堆
	h := &MinHeap{}
	heap.Init(h)

	// 将所有链表的头节点加入堆中（跳过空链表）
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	// 哑节点，方便构建结果链表
	dummy := &ListNode{Val: -1}
	cur := dummy
	// 当堆不为空时，不断弹出最小节点并加入结果链表
	for h.Len() > 0 {
		// 弹出当前最小的节点
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = cur.Next

		// 如果该节点还有下一个节点，则将下一个节点入堆
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	// 返回合并后的链表头节点
	return dummy.Next
}

func main() {

}
