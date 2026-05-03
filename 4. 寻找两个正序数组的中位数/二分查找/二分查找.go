package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 思路：
	// ● 中位数将数组分成左右两部分，且左边元素个数等于右边（或相差1）
	// ● 在较短数组上二分查找分割点 i，使得 nums1[i-1] <= nums2[j] 且 nums2[j-1] <= nums1[i]，其中 j = (m+n+1)/2 - i
	// ● 分割后，左边最大值为 max(nums1[i-1], nums2[j-1])，右边最小值为 min(nums1[i], nums2[j])，根据总长度奇偶计算中位数
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	// 左边需要包含的元素个数（奇数时左边多一个）
	totalLeft := (m + n + 1) / 2

	for left <= right {
		// 分割点 i 表示 nums1 中前 i 个元素在左半部分
		i := (left + right) / 2
		// 对应的 nums2 中前 j 个元素在左半部分
		j := totalLeft - i
		// 处理边界情况
		var leftMax1, leftMax2 int
		var rightMin1, rightMin2 int

		if i == 0 {
			leftMax1 = math.MinInt // 负无穷
		} else {
			leftMax1 = nums1[i-1]
		}
		if i == m {
			rightMin1 = math.MaxInt // 正无穷
		} else {
			rightMin1 = nums1[i]
		}
		if j == 0 {
			leftMax2 = math.MinInt // 负无穷
		} else {
			leftMax2 = nums2[j-1]
		}
		if j == n {
			rightMin2 = math.MaxInt // 正无穷
		} else {
			rightMin2 = nums2[j]
		}

		// 检查分割是否满足条件
		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			// 找到正确分割
			if (m+n)%2 == 1 {
				return float64(max(leftMax1, leftMax2))
			} else {
				return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2.0
			}
		} else if leftMax1 > rightMin2 {
			// i 太大，需要左移
			right = i - 1
		} else {
			// i 太小，需要右移
			left = i + 1
		}
	}

	return 0
}

func main() {

}
