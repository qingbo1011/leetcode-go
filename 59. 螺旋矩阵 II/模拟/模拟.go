package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	// 初始化res
	var res [][]int
	for i := 0; i < n; i++ { // 先把正方形框架定下来（因为是切片，不然后面直接赋值会out of index）
		res = append(res, make([]int, n))
	}
	// 定义一些变量
	count := 1                 // 赋值给螺旋矩阵的具体值
	i, j := 0, 0               // 遍历过程中具体的坐标
	startX, startY := 0, 0     // 每一圈遍历时，开始的坐标
	offset := 1                // 在每一圈的赋值中控制边界
	for k := 0; k < n/2; k++ { // 赋值需要转n/2圈
		for j = startY; j < n-offset; j++ { // 向右遍历，第1条边
			res[startX][j] = count
			count++
		}
		for i = startX; i < n-offset; i++ { // 向下遍历，第2条边
			res[i][j] = count // 第1条边遍历结束后，j已经在第2条边的位置了
			count++
		}
		for ; j > startY; j-- { // 向左遍历，第3条边（不指定j的值是因为，第1条边遍历结束后，j已经在指定位置了）
			res[i][j] = count
			count++
		}
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}
		// 4条边遍历完成后，一圈就转完了，更新相关数据
		startX++
		startY++
		offset++
		j++
	}
	if n%2 != 0 { // 如果n是奇数，那么转n/2圈后最中间的值还是空的，此时需要赋值
		i++ // 别忘了i++（因为每条边的遍历都是左闭右开嘛）
		res[i][j] = count
	}
	return res
}
