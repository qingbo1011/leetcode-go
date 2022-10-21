# 455. 分发饼干

[LeetCode 455. 分发饼干](https://leetcode.cn/problems/assign-cookies/)

## 贪心算法

这题是一道简单题，运用一下贪心算法很容易就能想出思路（排序+双指针+go）：

- 首先将`g`和`s`进行升序排序（自然顺序）
- 接下来使用`i`、`j`指针分别`g[0]`和`s[0]`。很容易理解：当`s[j]`饼干可以满足`g[i]`孩子时，`ans++`，i，j都继续向后遍历；反之当`s[j]`饼干不能满足`g[i]`孩子时，只是`j++`去寻找更大的饼干。

题目很简单直接看代码即可。

代码如下：

```go
func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] { // s[j]饼干可以满足g[i]孩子，ans++，i，j都继续向后遍历
			ans++
			i++
			j++
		} else { // 当前饼干不能满足该孩子，j继续向后遍历找到更大的饼干
			j++
		}
	}
	return ans
}
```

> 上面是为了方便看出思路所以这么写的if-else，可以优化成：
>
> ```go
> func findContentChildren(g []int, s []int) int {
> 	ans := 0
> 	sort.Ints(g)
> 	sort.Ints(s)
> 	i, j := 0, 0
> 	for i < len(g) && j < len(s) {
> 		if g[i] <= s[j] { // s[j]饼干可以满足g[i]孩子，ans++，i，j都继续向后遍历
> 			ans++
> 			i++
> 		}
> 		j++
> 	}
> 	return ans
> }
> ```















