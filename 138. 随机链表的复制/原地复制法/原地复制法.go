package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 第一步：插入复制节点
	for cur := head; cur != nil; cur = cur.Next.Next {
		copy := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = copy
	}
	// 第二步：设置 random 指针
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	// 第三步：拆分链表
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
	}
	return newHead
}

func main() {

}
