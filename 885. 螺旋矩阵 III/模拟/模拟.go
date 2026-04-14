package main

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	total := rows * cols
	res := make([][]int, 0, total)
	// 方向顺序：东、南、西、北
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := rStart, cStart
	res = append(res, []int{r, c})
	if total == 1 {
		return res
	}
	step := 1 // 当前步长
	dirIndex := 0
	for len(res) < total {
		// 每个方向走 step 步
		for i := 0; i < 2; i++ { // 每个步长连续两个方向
			d := dirs[dirIndex]
			for j := 0; j < step; j++ {
				r += d[0]
				c += d[1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					res = append(res, []int{r, c})
				}
			}
			// 按顺时针顺序切换下一个方向。
			dirIndex = (dirIndex + 1) % 4
		}
		step++ // 每两个方向后步长增加
	}
	return res
}

func main() {

}
