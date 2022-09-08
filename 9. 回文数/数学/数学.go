package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	// 特殊情况处理
	if x >= 0 && x < 10 { // x是个位数，肯定是回文
		return true
	}
	if x < 0 || x%10 == 0 { // x为负数，不可能为回文数; x最后一位为0，也显然不是回文数
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}
	// x为奇数：判断x == revertedNumber/10；x为偶数，判断：x == revertedNumber
	return x == revertedNumber || x == revertedNumber/10
}
