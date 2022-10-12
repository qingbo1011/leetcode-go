package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		// 双指针就位
		pre := temp.Next
		next := temp.Next.Next
		// 利用temp、pre、next这三个指针开始交换pre和next这两个节点
		temp.Next = next
		pre.Next = next.Next
		next.Next = pre
		// 交换完成，更新temp继续循环
		temp = pre
	}
	return dummy.Next
}
