package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int, 0) // key为nums[i]，value为i（即索引）
	for i, num := range nums {
		if value, ok := m[target-num]; ok {
			res = append(res, value, i)
			break
		}
		m[num] = i
	}
	return res
}
