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
	// 映射原节点->新节点
	m := make(map[*Node]*Node)
	// 第一遍：创建所有新节点，并建立映射
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	// 第二遍：设置next和random
	for cur := head; cur != nil; cur = cur.Next {
		// m[cur]是当前原节点cur对应的新节点
		// m[cur.Next]是原节点cur.Next对应的新节点
		// 将新节点的next指向原节点下一个节点对应的新节点，从而复制了链表结构。
		// Random同理
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}

	return m[head]
}

func main() {

}
