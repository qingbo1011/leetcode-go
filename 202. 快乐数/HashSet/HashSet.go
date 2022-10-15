package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))

}

func isHappy(n int) bool {
	set := make(map[int]struct{}, 0)
	for {
		if n == 1 {
			return true
		}
		n = getSum(n)
		if _, ok := set[n]; ok { // 出现了重复，说明陷入无限循环了
			return false
		}
		set[n] = struct{}{}
	}
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum = sum + (n%10)*(n%10) // n最后一位进行平方
		n = n / 10                // n去除最后一位
	}
	return sum
}
