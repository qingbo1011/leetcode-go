package main

func main() {

}

func maxProfit(prices []int) int {
	stack := make([]int, 0)
	ans := 0
	for _, price := range prices {
		for len(stack) > 0 && price < stack[len(stack)-1] {
			if len(stack) >= 2 { // 保证能出现低-高-低结构
				ans = max(ans, stack[len(stack)-1]-stack[0])
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, price)
	}
	if len(stack) >= 2 { // 遍历结束后，栈内还有2个或以上元素
		ans = max(ans, stack[len(stack)-1]-stack[0])
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
