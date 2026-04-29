package main

import "container/heap"

// Price 定义堆中存储的元素：包含价格和时间戳
// 注意：因为同一个时间戳的价格可能被覆盖，存储时间戳用于判断该元素是否过期
type Price struct {
	price     int
	timestamp int
}

// ---------- 最大堆实现 ----------
// MaxHeap 类型是一个切片，元素类型为 Price，通过实现 heap.Interface 来构建最大堆
type MaxHeap []Price

// Len 返回堆的元素个数
func (h MaxHeap) Len() int { return len(h) }

// Less 定义比较规则：由于需要最大堆，所以当 h[i].price > h[j].price 时返回 true
func (h MaxHeap) Less(i, j int) bool { return h[i].price > h[j].price }

// Swap 交换两个元素的位置
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 向堆中添加元素（注意参数为 interface{}，需要类型断言为 Price）
// 该方法由 container/heap 包自动调用，通常不直接调用
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Price))
}

// Pop 从堆中移除并返回最后一个元素（注意：实际堆顶元素由 heap.Pop 处理）
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- 最小堆实现 ----------
// MinHeap 类型同样实现 heap.Interface，构建最小堆
type MinHeap []Price

func (h MinHeap) Len() int { return len(h) }

// 最小堆：当 h[i].price < h[j].price 时返回 true
func (h MinHeap) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Price)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- StockPrice 主结构 ----------
type StockPrice struct {
	prices           map[int]int // 映射：时间戳 -> 最新价格（即当前有效价格）
	maxHeap          MaxHeap     // 最大堆，用于快速获取最高价格（可能包含过期数据）
	minHeap          MinHeap     // 最小堆，用于快速获取最低价格（可能包含过期数据）
	currentTimestamp int         // 当前最大时间戳（最新时间）
	currentPrice     int         // 最新时间戳对应的价格
}

// Constructor 初始化 StockPrice 对象
func Constructor() StockPrice {
	return StockPrice{
		prices:           make(map[int]int),
		maxHeap:          MaxHeap{},
		minHeap:          MinHeap{},
		currentTimestamp: 0,
		currentPrice:     0,
	}
}

// Update 更新某个时间戳的价格
func (this *StockPrice) Update(timestamp int, price int) {
	// 1. 更新哈希表，记录该时间戳的最新价格
	this.prices[timestamp] = price
	// 2. 将新价格同时推入最大堆和最小堆（堆中储存 price 和 timestamp）
	heap.Push(&this.maxHeap, Price{price, timestamp})
	heap.Push(&this.minHeap, Price{price, timestamp})
	// 3. 如果当前时间戳大于等于已知最新时间戳，则更新 currentTimestamp 和 currentPrice
	if timestamp >= this.currentTimestamp {
		this.currentTimestamp = timestamp
		this.currentPrice = price
	}
}

// Current 返回最新股票价格
func (this *StockPrice) Current() int {
	return this.currentPrice
}

// Maximum 返回当前存在的股票中的最高价格
// 使用最大堆，但堆中可能存在过期数据（某个时间戳的价格已被更新）
// 因此采用“懒删除”策略：循环检查堆顶元素，如果其价格与哈希表中该时间戳的当前价格不一致，
// 则说明该元素已过期，将其弹出，直到堆顶有效为止。
func (this *StockPrice) Maximum() int {
	for {
		top := this.maxHeap[0] // 获取堆顶（最大价格）
		if this.prices[top.timestamp] == top.price {
			return top.price // 有效，返回价格
		}
		heap.Pop(&this.maxHeap) // 过期，弹出并继续检查
	}
}

// Minimum 返回当前存在的股票中的最低价格
// 同理使用最小堆，采用懒删除策略
func (this *StockPrice) Minimum() int {
	for {
		top := this.minHeap[0]
		if this.prices[top.timestamp] == top.price {
			return top.price
		}
		heap.Pop(&this.minHeap)
	}
}

func main() {

}
