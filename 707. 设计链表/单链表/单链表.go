package main

func main() {

}

type MyLinkedList struct {
	dummy *Node
	size  int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &Node{}, size: 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size { // index不合法
		return -1
	}
	cur := l.dummy
	for i := 0; i <= index; i++ { // 因为cur是从哨兵节点开始的
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size { // 如果index大于链表长度，则不会插入节点
		return
	}
	if index < 0 { // 如果index小于0，则在头部插入节点
		index = 0
	}
	// 接下来就是链表的添加操作了
	pre := l.dummy
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	pre.Next = &Node{Next: cur, Val: val}
	l.size++ // 别忘了链表size+1
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size { // index不合法
		return
	}
	// 开始删除节点
	pre := l.dummy
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	l.size-- // 别忘了链表size-1

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
