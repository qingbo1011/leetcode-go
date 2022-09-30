# 704. 二分查找

[LeetCode 704. 二分查找](https://leetcode.cn/problems/binary-search/)

## 二分查找

这是一道经典的二分查找题目。

二分查找的思路很简单，但是代码写起来却很容易犯错。主要体现在两个方面：

- 例如到底是 `for left < right` 还是 `for left <= right`
- 到底是`right = middle`呢，还是要`right = middle - 1`呢

区分上面两种方法，主要是看区间是怎么定义的。写二分法，区间的定义一般为两种，**左闭右闭**即`[left, right]`，或者左闭右开即[left, right)。

这里我们给出**左闭右闭**的区间，即：

- `for left <= right`
- `left = middle + 1`
- `right = middle - 1`

> 如果是左闭右开即`[left, right)`：
>
> - `for left < right`
> - `left = middle + 1`
> - `right = middle`

代码如下：

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
```



































