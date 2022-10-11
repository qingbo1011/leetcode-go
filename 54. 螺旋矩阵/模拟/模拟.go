package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	m, n := len(matrix), len(matrix[0]) // m行n列
	// 初始化一些值
	left, right := 0, n-1
	top, bottom := 0, m-1
	count := 0
	// 开始遍历
	for count < m*n { // 一共遍历m*n次
		for i := left; i <= right; i++ { // 第1条边，从左到右
			res = append(res, matrix[top][i])
			count++
		}
		if count >= m*n {
			break
		}
		top++
		for i := top; i <= bottom; i++ { // 第2条边，从上到下
			res = append(res, matrix[i][right])
			count++
		}
		if count >= m*n {
			break
		}
		right--
		for i := right; i >= left; i-- { // 第3条边，从右到左
			res = append(res, matrix[bottom][i])
			count++
		}
		if count >= m*n {
			break
		}
		bottom--
		for i := bottom; i >= top; i-- { // 第4条边，从下到上
			res = append(res, matrix[i][left])
			count++
		}
		left++
	}
	return res
}
