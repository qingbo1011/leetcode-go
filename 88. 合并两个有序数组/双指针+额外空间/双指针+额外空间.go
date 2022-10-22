package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m || j < n {
		// 考虑i或j已经遍历到末尾的情况
		if i >= m {
			temp = append(temp, nums2[j:]...) // 这里可以直接把剩下的全部append了
			break
		}
		if j >= n {
			temp = append(temp, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	copy(nums1, temp)
}
