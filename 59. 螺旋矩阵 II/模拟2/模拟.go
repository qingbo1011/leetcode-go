package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	// 初始化res
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	// 定义一些初始值
	left, right := 0, n-1
	top, bottom := 0, n-1
	count := 1
	for count <= n*n {
		for i := left; i <= right; i++ { // 第1条边，从左到右
			res[top][i] = count
			count++
		}
		top++                            // 准备遍历下一条边
		for i := top; i <= bottom; i++ { // 第2条边，从上到下
			res[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- { // 第3条边，从右到左
			res[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- { // 第4条边，从下到上
			res[i][left] = count
			count++
		}
		left++
	}
	return res
}
