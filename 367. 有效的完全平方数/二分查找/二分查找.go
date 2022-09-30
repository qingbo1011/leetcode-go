package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(5))
}

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 1, num/2
	for left < right {
		middle := (left + right + 1) >> 1
		if middle*middle == num {
			return true
		} else if middle*middle < num { // 更新left
			left = middle
		} else { // 更新right
			right = middle - 1
		}
	}
	return false
}
