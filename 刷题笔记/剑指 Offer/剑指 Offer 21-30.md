# 剑指 Offer 29. 顺时针打印矩阵

[剑指 Offer 29. 顺时针打印矩阵](https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

这题和[LeetCode 54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)是一样的，只不过这题m和n可能为0，需要考虑m为0的情况。

具体分析这里就不写了，在[LeetCode 54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)的题解中有写过。

代码如下：

```go
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	m := len(matrix) // m行
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0]) // n列
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
```



