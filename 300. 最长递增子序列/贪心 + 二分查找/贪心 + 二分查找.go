package main

import "fmt"

func lengthOfLIS(nums []int) int {
	// 维护一个数组tails，其中tails[i] 表示长度为i+1的递增子序列的最小末尾元素。
	tails := make([]int, 0, len(nums))
	//遍历 nums 中的每个数 x，将x插入或替换到tails的合适位置(二分查找)：
	//● 如果 x 大于 tails 中所有元素，则追加到末尾，表示可以构造更长的递增子序列。
	//● 用x替换tails中第一个大于等于x的值的位置，使得该长度的递增子序列末尾变得更小，从而更有利于后续元素的扩展。
	for _, x := range nums {
		// 二分查找合适位置
		left, right := 0, len(tails)-1
		for left <= right {
			mid := left + (right-left)/2
			if tails[mid] < x {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left == len(tails) {
			tails = append(tails, x)
		} else {
			tails[left] = x
		}
	}
	// 最终tails的长度即为最长递增子序列的长度
	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
