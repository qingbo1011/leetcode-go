package main

type StockSpanner struct {
	stack [][2]int // [price, span]
}

func Constructor() StockSpanner {
	return StockSpanner{stack: [][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	// 维护一个单调递减栈（严格递减），栈中每个元素存储 (price, span)，其中 span 表示以该价格为顶点的连续小于等于它的天数（即从该天开始往前的跨度）。
	//当新价格 price 到来时：
	//● 初始化 total = 1（至少包括今天）。
	//● 循环弹出栈顶所有价格 ≤ price 的元素，并将它们的 span 累加到 total。
	//● 最后将 (price, total) 压入栈。
	//● 返回 total。
	span := 1
	// 弹出所有价格 <= 当前价格的栈顶元素，累加它们的 span
	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		span = span + this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	// 入栈当前价格和累计跨度
	this.stack = append(this.stack, [2]int{price, span})
	return span
}

func main() {

}
