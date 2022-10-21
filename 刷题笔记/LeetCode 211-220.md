# 217. 存在重复元素

[LeetCode 217. 存在重复元素](https://leetcode.cn/problems/contains-duplicate/)

## HashSet

LeetCode上的简单题。比较纯真😊。直接HashSet。没什么好说的。

代码如下：

```go
func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := set[num]; ok { // 出现了重复数据
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
```

> 复杂度分析：
>
> - 时间复杂度：O(N)，其中N为数组的长度。
> - 空间复杂度：O(N)，其中N为数组的长度。



