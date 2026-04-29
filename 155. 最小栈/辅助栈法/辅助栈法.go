package main

type MinStack struct {
	stack    []int // 主栈，存储所有元素
	minStack []int // 辅助栈，存储当前最小值序列
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	// 设计思路（辅助栈法）：
	// 使用主栈 stack 正常存储所有元素。
	// 使用辅助栈 minStack 同步维护当前栈中的最小值：
	//	每次 push 时，将新元素与 minStack 栈顶比较，若新元素 ≤ 当前最小值，则也压入 minStack。
	//	每次 pop 时，若弹出的元素等于 minStack 栈顶，则 minStack 也弹出栈顶。
	//	这样 minStack 的栈顶始终是当前主栈中的最小值。
	// 注意：当新元素等于当前最小值时，也必须压入 minStack，否则弹出时会丢失最小值记录。

	// 主栈正常压入
	this.stack = append(this.stack, val)
	// 如果辅助栈为空 或者 新元素 ≤ 辅助栈顶（当前最小值），则压入辅助栈
	// 注意：当 val == 当前最小值时也必须压入，否则多个相同最小值弹出时会丢失记录
	if len(this.minStack) == 0 || val <= this.GetMin() {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// 获取栈顶元素
	top := this.Top()
	// 主栈弹出
	this.stack = this.stack[:len(this.stack)-1]
	// 如果弹出的元素等于辅助栈顶（即当前最小值），辅助栈也弹出
	if top == this.GetMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {

}
