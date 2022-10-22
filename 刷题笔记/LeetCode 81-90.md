# 88. 合并两个有序数组

[LeetCode 88. 合并两个有序数组](https://leetcode.cn/problems/merge-sorted-array/)

## 双指针+额外空间

这个没什么好说的。使用额外空间就很好写这题了。直接看代码吧。

代码如下：

```go
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
```

## 逆序双指针

> 所有玩家都全力向前冲刺, 却不知道向后才是胜利之门。——《头号玩家》

想要不使用额外空间，就需要使用**逆序双指针**了。

参考[官方题解](https://leetcode.cn/problems/merge-sorted-array/solution/he-bing-liang-ge-you-xu-shu-zu-by-leetco-rrb0/)

三个指针就位情况：`i, j, tail := m-1, n-1, m+n-1`

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221022111855.png)

代码如下：

```go
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
```





















