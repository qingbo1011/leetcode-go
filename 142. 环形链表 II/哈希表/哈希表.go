package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	p := head
	for p != nil {
		if _, ok := m[p]; ok { // 找到重复索引了
			return p
		}
		m[p] = struct{}{}
		p = p.Next
	}
	return nil
}
