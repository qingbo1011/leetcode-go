package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(7))
}

func mySqrt(x int) int {
	// 特殊值判断
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x/2
	for left < right {
		middle := (left + right + 1) >> 1
		if middle*middle == x {
			return middle
		} else if middle*middle < x { // 更新left（结果可能是middle）
			left = middle
		} else { // 更新right（结果不可能是right，因为都middle*middle大于x了）
			right = middle - 1
		}
	}
	return left
}
