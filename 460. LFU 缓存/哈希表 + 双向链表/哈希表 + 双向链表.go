package main

// Node 双向链表的节点，存储缓存的键值对以及使用频率和前后指针
type Node struct {
	key   int
	value int
	freq  int // 访问频率
	prev  *Node
	next  *Node
}

// LinkedList 双向链表，包含虚拟头尾节点，用于管理同一频率的所有节点（按访问时间排序，头部最近使用，尾部最久未使用）
type LinkedList struct {
	head *Node // 虚拟头节点（不存储实际数据）
	tail *Node // 虚拟尾节点（不存储实际数据）
}

// newLinkedList 创建一个空的双向链表，初始化虚拟头尾节点并互相连接
func newLinkedList() *LinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LinkedList{head: head, tail: tail}
}

// addToHead 将节点插入链表头部（标记为最近使用）
// 插入位置：head 和原 head.next 之间
func (l *LinkedList) addToHead(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

// removeNode 从链表中删除指定节点（节点必须存在）
func (l *LinkedList) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeTail 删除链表尾部节点（最久未使用）并返回该节点
// 若链表为空（仅虚拟头尾），返回 nil
func (l *LinkedList) removeTail() *Node {
	if l.tail.prev == l.head {
		return nil
	}
	node := l.tail.prev
	l.removeNode(node)
	return node
}

// isEmpty 判断链表是否为空（即没有实际存储数据的节点）
func (l *LinkedList) isEmpty() bool {
	return l.head.next == l.tail
}

type LFUCache struct {
	capacity   int                 // 缓存容量
	minFreq    int                 // 当前缓存中最小的访问频率
	keyToNode  map[int]*Node       // 键到节点的映射（O(1) 查找）
	freqToList map[int]*LinkedList // 频率到双向链表的映射（每个频率对应一个链表）
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		minFreq:    0,
		keyToNode:  make(map[int]*Node),
		freqToList: make(map[int]*LinkedList),
	}
}

// updateNode 更新节点的频率：
// 1. 从原频率链表中移除节点
// 2. 若移除后该链表为空且该频率等于 minFreq，则 minFreq++
// 3. 节点频率加1
// 4. 将节点插入到新频率链表的头部（最近使用位置）
func (this *LFUCache) updateNode(node *Node) {
	// 从原频率链表中移除
	list := this.freqToList[node.freq]
	list.removeNode(node)
	// 如果该链表变空且该频率是当前最小频率，则最小频率加1
	if list.isEmpty() && node.freq == this.minFreq {
		this.minFreq++
	}
	// 频率增加
	node.freq++
	// 获取或创建新频率对应的链表
	newList, ok := this.freqToList[node.freq]
	if !ok {
		newList = newLinkedList()
		this.freqToList[node.freq] = newList
	}
	// 将节点插入新链表的头部（作为最近使用）
	newList.addToHead(node)
}

// Get 获取键对应的值，若不存在返回 -1；存在则更新频率并返回值
func (this *LFUCache) Get(key int) int {
	node, ok := this.keyToNode[key]
	if !ok {
		return -1
	}
	this.updateNode(node)
	return node.value
}

// Put 插入或更新键值对，若超出容量则淘汰最不经常使用的节点（同频率淘汰最久未使用）
func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 如果键已存在，更新值并增加频率
	if node, ok := this.keyToNode[key]; ok {
		node.value = value
		this.updateNode(node)
		return
	}
	// 缓存已满，需淘汰一个节点
	if len(this.keyToNode) == this.capacity {
		// 从最小频率的链表中删除尾部节点（最久未使用）
		list := this.freqToList[this.minFreq]
		tail := list.removeTail()
		// 同时删除哈希表中的记录
		delete(this.keyToNode, tail.key)
	}
	// 创建新节点，频率为1
	newNode := &Node{key: key, value: value, freq: 1}
	this.keyToNode[key] = newNode
	// 获取或创建频率为1的链表
	list, ok := this.freqToList[1]
	if !ok {
		list = newLinkedList()
		this.freqToList[1] = list
	}
	// 将新节点插入头部（最近使用）
	list.addToHead(newNode)
	// 新插入的节点频率为1，因此最小频率更新为1
	this.minFreq = 1
}

func main() {

}
