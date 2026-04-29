package main

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache
// 使用双向链表维护缓存中键值对的访问顺序：头部（head）为最近使用，尾部（tail）为最久未使用。
// 使用哈希表（map[int]*Node）存储键到链表节点的映射，保证 O(1) 时间定位节点。
// head和tail是虚拟节点（dummy node），它们不存储任何实际的键值对（key 和 value 字段无意义），仅作为链表的边界标志。
type LRUCache struct {
	capacity int
	cache    map[int]*Node // key -> node 映射
	head     *Node         // 虚拟头节点
	tail     *Node         // 虚拟尾节点
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
	// 初始化双向链表：head <-> tail
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// 将节点移动到链表头部（最近使用）
func (this *LRUCache) moveToHead(node *Node) {
	// 先断开原位置
	node.prev.next = node.next
	node.next.prev = node.prev
	// 插入头部
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

// 在头部插入新节点
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 删除尾部节点（最久未使用）
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	node.prev.next = this.tail
	this.tail.prev = node.prev
	return node
}

func (this *LRUCache) Get(key int) int {
	// 若key存在，将对应节点移动到链表头部，并返回节点值；否则返回-1。
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 若key已存在，更新节点值并将节点移动到头部；
	if node, ok := this.cache[key]; ok {
		// 更新值并移动到头部
		node.value = value
		this.moveToHead(node)
		return
	}
	// 若key不存在，创建新节点插入头部，若超出容量则删除尾部节点，并删除哈希表中的记录。
	// 创建新节点
	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.addToHead(newNode)
	if len(this.cache) > this.capacity {
		// 移除尾部节点
		tailNode := this.removeTail()
		delete(this.cache, tailNode.key)
	}
}

func main() {

}
