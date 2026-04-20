package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap 自定义最大堆
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // 最大堆
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	// 初始化最大堆
	h := MaxHeap(stones)
	heap.Init(&h)

	// 循环直到堆中元素少于2
	for h.Len() >= 2 {
		x := heap.Pop(&h).(int) // 最重
		y := heap.Pop(&h).(int) // 次重
		if x > y {
			heap.Push(&h, x-y)
		}
	}
	// 如果最后堆长度为0，说明没有石头剩下
	if h.Len() == 0 {
		return 0
	}

	return h[0]
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
