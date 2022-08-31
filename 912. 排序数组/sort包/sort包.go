package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}
