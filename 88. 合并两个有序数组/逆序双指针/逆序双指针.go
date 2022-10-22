package main

import "fmt"

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[tail] = nums1[i]
				i--
				tail--
			} else {
				nums1[tail] = nums2[j]
				j--
				tail--
			}
		} else { // 考虑其中一个指针遍历结束遍历结束而另一个指针没有遍历结束的情况
			if i < 0 {
				nums1[tail] = nums2[j]
				j--
				tail--
			} else if j < 0 { // 注意这里要else if而不是两个单独的if
				nums1[tail] = nums1[i]
				i--
				tail--
			}
		}
	}
}
