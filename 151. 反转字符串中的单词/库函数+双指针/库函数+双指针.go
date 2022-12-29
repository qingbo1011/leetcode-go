package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))

}

func reverseWords(s string) string {
	res := ""
	fields := strings.Fields(s)
	left, right := 0, len(fields)-1
	for left < right {
		fields[left], fields[right] = fields[right], fields[left]
		left++
		right--
	}
	for i := 0; i < len(fields)-1; i++ {
		res = res + fields[i] + " "
	}
	res = res + fields[len(fields)-1] // 加上最后一个
	return res
}
