package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x >= 0 && x < 10 { // x是个位数，肯定是回文
		return true
	}
	if x < 0 || x%10 == 0 { // x为负数，不可能为回文数; x最后一位为0，也显然不是回文数
		return false
	}
	str := strconv.Itoa(x)
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
