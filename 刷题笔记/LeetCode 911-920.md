# 912. 排序数组

[LeetCode 912. 排序数组](https://leetcode.cn/problems/sort-an-array/)

## 快速排序

考查各大排序算法，参考笔记：**[Go排序](https://www.qingbo1011.top/2022/04/28/Go%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/)**

快排代码如下：

```go
func sortArray(nums []int) []int {
    return quicksort(nums)
}

func quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	left, right := 0, len(s)-1
	pivot := 0 // pivot可以在[0,len(s)-1]之间随意选择
	s[pivot], s[right] = s[right], s[pivot]
	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}
	s[left], s[right] = s[right], s[left]

	quicksort(s[:left])
	quicksort(s[left+1:])

	return s
}
```

> 手写的快速排序，结果为：
>
> - 执行用时：1084 ms, 在所有 Go 提交中击败了7.30%的用户
> - 内存消耗：12.7 MB, 在所有 Go 提交中击败了29.23%的用户
>
> 如果用Go自带的排序方法：
>
> ```go
> import (
> 	"sort"
> )
> 
> func sortArray(nums []int) []int {
>     sort.Ints(nums)
>     return nums
> }
> ```
>
> 效率还高一些（笔记Go官方进行了封装，可以看看[进阶：pdqsort](https://www.qingbo1011.top/2022/04/28/Go%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/#%E8%BF%9B%E9%98%B6pdqsort)的笔记）















